// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, number } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer } from '@commonground/design-system'
import DirectoryDetailView from './components/DirectoryDetailView'
import DrawerHeader from './components/DrawerHeader'
import { StyledDrawer } from './index.styles'

const DirectoryDetailPage = ({ service, parentUrl }) => {
  const history = useHistory()
  const { serviceName } = useParams()

  const navigateToParentUrl = () => {
    history.push(parentUrl)
  }

  return (
    <StyledDrawer
      id="directoryDetail"
      noMask
      afterHide={navigateToParentUrl}
      closeHandler={navigateToParentUrl}
    >
      {service ? (
        <DrawerHeader service={service} />
      ) : (
        <Drawer.Header
          as="header"
          title={serviceName}
          closeButtonLabel="Close"
        />
      )}

      <Drawer.Content>
        {service ? (
          <DirectoryDetailView service={service} />
        ) : (
          <Alert variant="error" data-testid="error-message">
            {`Kan de service '${serviceName}' niet vinden.`}
          </Alert>
        )}
      </Drawer.Content>
    </StyledDrawer>
  )
}

DirectoryDetailPage.propTypes = {
  service: shape({
    apiType: string,
    contactEmailAddress: string,
    documentationUrl: string,
    name: string.isRequired,
    organization: string.isRequired,
    status: string.isRequired,
    oneTimeCosts: number,
    monthlyCosts: number,
    requestCosts: number,
  }),
  parentUrl: string,
}

DirectoryDetailPage.defaultProps = {
  parentUrl: '/directory',
}

export default DirectoryDetailPage
