import { Component, OnInit, OnDestroy } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Store } from '@ngrx/store';
import { Subject, combineLatest } from 'rxjs';
import { filter, map, pluck, takeUntil } from 'rxjs/operators';
import { identity, some } from 'lodash/fp';

import { NgrxStateAtom } from 'app/ngrx.reducers';
import { routeParams } from 'app/route.selectors';
import { Regex } from 'app/helpers/auth/regex';
import { EntityStatus, loading } from 'app/entities/entities';
import {
  getStatus, updateStatus, projectFromRoute
} from 'app/entities/projects/project.selectors';
import { Project } from 'app/entities/projects/project.model';
import { GetProject, UpdateProject } from 'app/entities/projects/project.actions';
import { GetRulesForProject, DeleteRule } from 'app/entities/rules/rule.actions';
import { Rule, RuleStatus } from 'app/entities/rules/rule.model';
import {
  allRules
} from 'app/entities/rules/rule.selectors';

export type ProjectTabName = 'rules' | 'details';

@Component({
  selector: 'app-project-details',
  templateUrl: './project-details.component.html',
  styleUrls: ['./project-details.component.scss']
})

export class ProjectDetailsComponent implements OnInit, OnDestroy {
  public project: Project;
  public projectForm: FormGroup;
  public saveSuccessful = false;
  public isChefManaged = false;
  public rules: Rule[] = [];
  public selectedTab: ProjectTabName = 'rules';
  public ruleToDelete: Rule;
  public deleteModalVisible = false;
  public createModalVisible = false;
  public createProjectForm: FormGroup;
  public creatingProject = false;
  // isLoading represents the initial load as well as subsequent updates in progress.
  public isLoading = true;
  public saving = false;
  private isDestroyed = new Subject<boolean>();

  constructor(
    private fb: FormBuilder,
    private store: Store<NgrxStateAtom>
  ) { }

  ngOnInit(): void {

    this.projectForm = this.fb.group({
      // Must stay in sync with error checks in project-details.component.html
      name: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]]
    });

    combineLatest(
      this.store.select(getStatus),
      this.store.select(updateStatus)
    ).pipe(
      takeUntil(this.isDestroyed),
      map(([gStatus, uStatus]) => {
        this.isLoading =
          (gStatus !== EntityStatus.loadingSuccess) ||
          (uStatus === EntityStatus.loading);
      })
    ).subscribe();

    this.store.select(projectFromRoute).pipe(
      filter(identity),
      takeUntil(this.isDestroyed),
      map((state) => {
        this.project = <Project>Object.assign({}, state);
        this.store.dispatch(new GetRulesForProject({ project_id: this.project.id }));
        this.store.select(allRules).subscribe((rules) => {
          this.rules = rules;
        });
        this.isChefManaged = this.project.type === 'CHEF_MANAGED';
        this.projectForm.controls['name'].setValue(this.project.name);
      })).subscribe();

    this.store.select(routeParams).pipe(
      pluck('id'),
      filter(identity),
      takeUntil(this.isDestroyed))
      .subscribe((id: string) => {
        this.store.dispatch(new GetProject({ id }));
      });
  }

  ngOnDestroy() {
    this.isDestroyed.next(true);
    this.isDestroyed.complete();
  }

  keyPressed() {
    this.saveSuccessful = false;
  }

  onTabChange(event: { target: { value: ProjectTabName } }) {
    this.selectedTab = event.target.value;
  }

  showTab(tabName: ProjectTabName): boolean {
    return this.selectedTab === tabName;
  }

  showFirstRuleMessage(): boolean {
    return this.rules.length === 0;
  }

  showRulesTable(): boolean {
    return this.rules.length > 0;
  }

  closeDeleteModal(): void {
    this.deleteModalVisible = false;
  }

  startRuleDelete(rule: Rule): void {
    this.deleteModalVisible = true;
    this.ruleToDelete = rule;
  }

  deleteRule(): void {
    this.store.dispatch(new DeleteRule({
      project_id: this.ruleToDelete.project_id,
      id: this.ruleToDelete.id
    }));
    this.closeDeleteModal();
  }

  getEditStatus(rule: Rule): string {
    return rule.status === 'STAGED' ? 'Edits pending' : 'Applied';
  }

  showDeleteRule(): boolean {
    return true; // TODO: return false when *project* status is "updating..."
  }

  showProjectLink(): boolean {
    const statusPropertyName = 'status';
    const ruleStatus: RuleStatus = 'STAGED';
    return some([statusPropertyName, ruleStatus], this.rules);
  }

  saveProject() {
    this.saveSuccessful = false;
    this.saving = true;
    this.store.dispatch(new UpdateProject({
      id: this.project.id,
      name: this.projectForm.controls['name'].value.trim()
    }));

    const pendingSave = new Subject<boolean>();
    this.store.select(updateStatus).pipe(
      filter(identity),
      takeUntil(pendingSave))
      .subscribe((state) => {
        if (!loading(state)) {
          pendingSave.next(true);
          pendingSave.complete();
          this.saving = false;
          this.saveSuccessful = (state === EntityStatus.loadingSuccess);
        }
      });
  }
}
