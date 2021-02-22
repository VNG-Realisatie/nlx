// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

// The values to be used for testing in the review app environment are set in .gitlab/ci/deploy/review.yml

export const INWAY_SELF_ADDRESS =
  process.env.E2E_MANAGEMENT_UI_INWAY_SELF_ADDRESS ||
  'inway.organization-a.nlx.local:7913'
export const INWAY_VERSION =
  process.env.E2E_MANAGEMENT_UI_INWAY_VERSION || 'unknown'

export const DIRECTORY_ORGANIZATION_NAME =
  process.env.E2E_MANAGEMENT_UI_DIRECTORY_ORGANIZATION_NAME || 'Organization-B'
export const DIRECTORY_SERVICE_NAME =
  process.env.E2E_MANAGEMENT_UI_DIRECTORY_SERVICE_NAME || 'Petstore'
export const DIRECTORY_STATUS =
  process.env.E2E_MANAGEMENT_UI_DIRECTORY_STATUS || 'Beschikbaar'
export const DIRECTORY_API_SPECIFICATION_TYPE =
  process.env.E2E_MANAGEMENT_UI_DIRECTORY_API_SPECIFICATION_TYPE || 'OpenAPI2'
