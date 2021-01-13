// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { observer } from 'mobx-react'
import { object, func } from 'prop-types'
import { Table, ToasterContext, Collapsible } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import {
  DetailHeading,
  StyledCollapsibleBody,
} from '../../../../../components/DetailView'
import Amount from '../../../../../components/Amount'
import { IconKey } from '../../../../../icons'
import useStores from '../../../../../hooks/use-stores'
import useInterval from '../../../../../hooks/use-interval'
import IncomingAccessRequestRow from './IncomingAccessRequestRow'
import { StyledUpdateUiButton } from './index.styles'

const AccessRequestsSection = ({
  service,
  onApproveOrRejectCallbackHandler,
}) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const rootStore = useStores()

  useInterval(async () => {
    await rootStore.incomingAccessRequestsStore.fetchForService({
      name: service.name,
    })
  }, 3000)

  // start interval
  // haal incoming access requests op via apiClient (dus niet store)
  // vergelijk respons met huidige accessRequests property
  // if verschillend -> toon blauwe pil

  // blauwe pil
  // onclick -> voeg response van api toe aan store
  // verwijder BLAUWE KNOP

  const approveHandler = async (accessRequest) => {
    try {
      await accessRequest.approve()

      showToast({
        title: t('Access request approved'),
        body: t('Organization has access to service', {
          organizationName: accessRequest.organizationName,
          serviceName: accessRequest.serviceName,
        }),
        variant: 'success',
      })

      await onApproveOrRejectCallbackHandler()
    } catch (error) {
      showToast({
        title: t('Failed to approve access request'),
        body: t('Please try again'),
        variant: 'error',
      })
    }
  }

  const rejectHandler = async (accessRequest) => {
    try {
      await accessRequest.reject()

      showToast({
        title: t('Access request rejected'),
        body: t('Organization has been denied access to service', {
          organizationName: accessRequest.organizationName,
          serviceName: accessRequest.serviceName,
        }),
        variant: 'success',
      })

      await onApproveOrRejectCallbackHandler()
    } catch (error) {
      showToast({
        title: t('Failed to reject access request'),
        body: t('Please try again'),
        variant: 'error',
      })
    }
  }

  return (
    <Collapsible
      title={
        <DetailHeading data-testid="service-incoming-accessrequests">
          <IconKey />
          {t('Access requests')}
          {service.incomingAccessRequests.length > 0 ? (
            <Amount
              data-testid="service-incoming-accessrequests-amount-accented"
              value={service.incomingAccessRequests.length}
              isAccented
            />
          ) : (
            <Amount
              data-testid="service-incoming-accessrequests-amount"
              value={service.incomingAccessRequests.length}
            />
          )}
        </DetailHeading>
      }
      ariaLabel={t('Access requests')}
    >
      <StyledCollapsibleBody>
        {service.incomingAccessRequests.length ? (
          <Table data-testid="service-incoming-accessrequests-list">
            <tbody>
              {service.incomingAccessRequests.map((accessRequest) => (
                <IncomingAccessRequestRow
                  key={accessRequest.id}
                  accessRequest={accessRequest}
                  approveHandler={approveHandler}
                  rejectHandler={rejectHandler}
                />
              ))}
            </tbody>
          </Table>
        ) : (
          <small>{t('There are no access requests')}</small>
        )}

        <StyledUpdateUiButton>Nieuwe verzoeken</StyledUpdateUiButton>
      </StyledCollapsibleBody>
    </Collapsible>
  )
}

AccessRequestsSection.propTypes = {
  service: object,
  onApproveOrRejectCallbackHandler: func.isRequired,
}

export default observer(AccessRequestsSection)
