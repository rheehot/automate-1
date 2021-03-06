import { EntityState, EntityAdapter, createEntityAdapter } from '@ngrx/entity';
import { set, pipe } from 'lodash/fp';
import { EntityStatus } from 'app/entities/entities';
import { EnvironmentActionTypes, EnvironmentActions } from './environment.action';
import { Environment } from './environment.model';

export interface EnvironmentEntityState extends EntityState<Environment> {
  environmentsStatus: EntityStatus;
  getAllStatus: EntityStatus;
}

const GET_ALL_STATUS = 'getAllStatus';

export const environmentEntityAdapter: EntityAdapter<Environment> =
  createEntityAdapter<Environment>({
  selectId: (environment: Environment) => environment.name
});

export const EnvironmentEntityInitialState: EnvironmentEntityState =
  environmentEntityAdapter.getInitialState(<EnvironmentEntityState>{
  getAllStatus: EntityStatus.notLoaded
});

export function environmentEntityReducer(
  state: EnvironmentEntityState = EnvironmentEntityInitialState,
  action: EnvironmentActions): EnvironmentEntityState {

  switch (action.type) {
    case EnvironmentActionTypes.GET_ALL:
      return set(GET_ALL_STATUS, EntityStatus.loading, environmentEntityAdapter.removeAll(state));

    case EnvironmentActionTypes.GET_ALL_SUCCESS:
      return pipe(
        set(GET_ALL_STATUS, EntityStatus.loadingSuccess))
        (environmentEntityAdapter.setAll(action.payload.environments, state)) as
        EnvironmentEntityState;

    case EnvironmentActionTypes.GET_ALL_FAILURE:
      return set(GET_ALL_STATUS, EntityStatus.loadingFailure, state);

    default:
      return state;
  }
}

export const getEntityById = (id: string) =>
  (state: EnvironmentEntityState) => state.entities[id];
