describe('global projects filter', () => {
  before(() => {
    cy.login('/').then(() => {
      let admin = JSON.parse(localStorage.getItem('chef-automate-user'))
      cy.request({
          auth: { bearer: admin.id_token },
          method: 'POST',
          url: `/apis/iam/v2beta/projects`,
          failOnStatusCode: false,
          body: {
            id: "cypress-project-1",
            name: "Cypress Project One"
          }
      }).then((response) => {
        // if user already exists, move on
        expect([200, 409]).to.include(response.status)
        })
      cy.request({
          auth: { bearer: admin.id_token },
          method: 'POST',
          url: `/apis/iam/v2beta/projects`,
          failOnStatusCode: false,
          body: {
            id: "cypress-project-2",
            name: "Cypress Project Two"
          }
      }).then((response) => {
        expect([200, 409]).to.include(response.status)
      })
    })
  })

  after(() => {
    let admin = JSON.parse(localStorage.getItem('chef-automate-user'))
    cy.cleanupProjects(admin.id_token)
  }) 

  it('shows allowed projects for selection', () => {
    cy.visit('/settings')
    // hide modal unrelated to test flow
    cy.get('app-welcome-modal').invoke('hide')

    cy.get('chef-sidebar')
      .should('have.attr', 'minor-version')
      .then((version) => {
        if (version === 'v1') {
          cy.get('[data-cy=projects-filter-button]').click()
          
          const allowedProjects = ['cypress-project-1', 'cypress-project-2', '(unassigned)'];
          allowedProjects.forEach(project => {
            cy.get('[data-cy=projects-filter-dropdown]').contains(project)
          })
        }
      })
  })
})

