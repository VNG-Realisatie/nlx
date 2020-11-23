// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, array, bool, func, shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Link, useLocation } from 'react-router-dom'

import EditButton from '../../../../components/EditButton'
import {
  DetailHeadingLight,
  SectionGroup,
} from '../../../../components/DetailView'
import { IconHidden, IconVisible } from '../../../../icons'
import { showServiceVisibilityAlert } from '../../../../components/ServiceVisibilityAlert'
import InwaysSection from './InwaysSection'
import AccessRequestsSection from './AccessRequestsSection'
import AccessGrantSection from './AccessGrantSection'
import {
  StyledActionsBar,
  StyledRemoveButton,
  StyledServiceVisibilityAlert,
} from './index.styles'

const ServiceDetailView = ({ service, removeHandler }) => {
  const { internal, inways } = service
  const { t } = useTranslation()
  const location = useLocation()

  const handleRemove = () => {
    if (window.confirm(t('Do you want to remove the service?'))) {
      removeHandler()
    }
  }

  return (
    <>
      {showServiceVisibilityAlert({ internal, inways }) ? (
        <StyledServiceVisibilityAlert />
      ) : null}

      <StyledActionsBar>
        <EditButton
          as={Link}
          to={`${location.pathname}/edit-service`}
          title={t('Edit service')}
        />
        <StyledRemoveButton
          title={t('Remove service')}
          onClick={handleRemove}
        />
      </StyledActionsBar>

      <SectionGroup>
        <DetailHeadingLight>
          {internal ? (
            <>
              <IconHidden />
              {t('Not visible in central directory')}
            </>
          ) : (
            <>
              <IconVisible />
              {t('Published in central directory')}
            </>
          )}
        </DetailHeadingLight>

        <InwaysSection inways={inways} />

        <AccessRequestsSection
          accessRequests={service.incomingAccessRequests}
          onApproveOrRejectCallbackHandler={() => {
            service.fetch()

            service.update({
              incomingAccessRequestsCount:
                service.incomingAccessRequestsCount - 1,
            })
          }}
        />

        <AccessGrantSection accessGrants={service.accessGrants} />
      </SectionGroup>
    </>
  )
}

ServiceDetailView.propTypes = {
  service: shape({
    endpointURL: string,
    documentationURL: string,
    apiSpecificationURL: string,
    internal: bool.isRequired,
    techSupportContact: string,
    publicSupportContact: string,
    inways: arrayOf(string),
    incomingAccessRequests: array,
  }).isRequired,
  removeHandler: func.isRequired,
}

ServiceDetailView.defaultProps = {}

export default observer(ServiceDetailView)
