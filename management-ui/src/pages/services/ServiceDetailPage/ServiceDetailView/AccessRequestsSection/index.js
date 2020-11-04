// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { observer } from 'mobx-react'
import { array } from 'prop-types'
import { Table, ToasterContext, Collapsible } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import {
  DetailHeading,
  StyledCollapsibleBody,
  StyledCollapsibleEmptyBody,
} from '../../../../../components/DetailView'
import Amount from '../../../../../components/Amount'
import { IconKey } from '../../../../../icons'
import IncomingAccessRequestRow from './IncomingAccessRequestRow'

const AccessRequestsSection = ({ accessRequests }) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)

  const approveHandler = async (accessRequest) => {
    await accessRequest.approve()

    if (!accessRequest.error) {
      showToast({
        title: t('Access request approved'),
        body: t('Organization has access to service', {
          organizationName: accessRequest.organizationName,
          serviceName: accessRequest.serviceName,
        }),
        variant: 'success',
      })
    } else {
      showToast({
        title: t('Failed to approve access request'),
        body: t('Please try again'),
        variant: 'error',
      })
      console.error(accessRequest.error)
    }
  }

  const rejectHandler = async (accessRequest) => {
    await accessRequest.reject()

    if (!accessRequest.error) {
      showToast({
        title: t('Access request rejected'),
        body: t('Organization has been denied access to service', {
          organizationName: accessRequest.organizationName,
          serviceName: accessRequest.serviceName,
        }),
        variant: 'success',
      })
    } else {
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
          {accessRequests.length > 0 ? (
            <Amount
              data-testid="service-incoming-accessrequests-amount-accented"
              value={accessRequests.length}
              isAccented
            />
          ) : (
            <Amount
              data-testid="service-incoming-accessrequests-amount"
              value={accessRequests.length}
            />
          )}
        </DetailHeading>
      }
      ariaLabel={t('Access requests')}
    >
      <StyledCollapsibleBody>
        <Table data-testid="service-incoming-accessrequests-list">
          <tbody>
            {accessRequests.length ? (
              accessRequests.map((accessRequest) => (
                <IncomingAccessRequestRow
                  key={accessRequest.id}
                  accessRequest={accessRequest}
                  approveHandler={approveHandler}
                  rejectHandler={rejectHandler}
                />
              ))
            ) : (
              <Table.Tr data-testid="service-no-incoming-accessrequests">
                <Table.Td>
                  <StyledCollapsibleEmptyBody>
                    {t('There are no access requests')}
                  </StyledCollapsibleEmptyBody>
                </Table.Td>
              </Table.Tr>
            )}
          </tbody>
        </Table>
      </StyledCollapsibleBody>
    </Collapsible>
  )
}

AccessRequestsSection.propTypes = {
  accessRequests: array,
}

AccessRequestsSection.defaultProps = {
  accessRequests: [],
}

export default observer(AccessRequestsSection)
