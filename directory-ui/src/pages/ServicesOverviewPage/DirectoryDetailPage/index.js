// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useEffect } from 'react'
import { shape, string } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import {
  Alert,
  Drawer,
  withDrawerStack,
  useDrawerStack,
} from '@commonground/design-system'
import DirectoryDetailView from './components/DirectoryDetailView'
import DrawerHeader from './components/DrawerHeader'
import { StyledStackedDrawer } from './index.styles'

const DirectoryDetailPage = ({ service, parentUrl }) => {
  const history = useHistory()
  const { organizationName, serviceName } = useParams()
  const { showDrawer } = useDrawerStack()

  const navigateToParentUrl = () => {
    history.push(parentUrl)
  }

  useEffect(() => {
    showDrawer('directoryDetail')
  }, [service]) // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <StyledStackedDrawer
      id="directoryDetail"
      noMask
      afterHide={navigateToParentUrl}
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
            {`Failed to load the service ${organizationName}/${serviceName}`}
          </Alert>
        )}
      </Drawer.Content>
    </StyledStackedDrawer>
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
  }),
  parentUrl: string,
}

DirectoryDetailPage.defaultProps = {
  parentUrl: '/directory',
}

export default withDrawerStack(DirectoryDetailPage)
