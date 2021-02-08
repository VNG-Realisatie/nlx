// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useContext } from 'react'
import { observer } from 'mobx-react'
import { array, bool, func } from 'prop-types'
import { Table, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { StyledCollapsibleBody } from '../../../../../../components/DetailView'
import UpdateUiButton from '../../UpdateUiButton'
import IncomingAccessRequestRow from './IncomingAccessRequestRow'
import { StyledSmall } from './index.styles'

const CollapsibleBody = ({
  accessRequests,
  showLoadIncomingDataButton,
  onClickLoadIncomingDataHandler,
  onApproveOrRejectCallbackHandler,
}) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)

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
    <StyledCollapsibleBody>
      {accessRequests.length ? (
        <Table data-testid="service-incoming-accessrequests-list">
          <tbody>
            {accessRequests.map((accessRequest) => (
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
        <StyledSmall>{t('There are no access requests')}</StyledSmall>
      )}

      {showLoadIncomingDataButton ? (
        <UpdateUiButton onClick={onClickLoadIncomingDataHandler}>
          {t('Show updates')}
        </UpdateUiButton>
      ) : null}
    </StyledCollapsibleBody>
  )
}

CollapsibleBody.propTypes = {
  accessRequests: array,
  showLoadIncomingDataButton: bool,
  onClickLoadIncomingDataHandler: func,
  onApproveOrRejectCallbackHandler: func,
}

const noop = () => {}

CollapsibleBody.defaultProps = {
  showLoadIncomingDataButton: false,
  onClickLoadIncomingDataHandler: noop,
  onApproveOrRejectCallbackHandler: noop,
}

export default observer(CollapsibleBody)
