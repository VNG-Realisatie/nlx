// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useContext } from 'react'
import { ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { array, bool, func } from 'prop-types'
import Table from '../../../../../../components/Table'
import { StyledCollapsibleBody } from '../../../../../../components/DetailView'
import UpdateUiButton from '../../UpdateUiButton'
import AccessGrantRow from './AccessGrantRow'
import { StyledSmall } from './index.styles'

const CollapsibleBody = ({
  accessGrants,
  showLoadIncomingDataButton,
  onClickLoadIncomingDataHandler,
}) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)

  const revokeHandler = async (accessGrant) => {
    try {
      await accessGrant.revoke()

      showToast({
        title: t('Access revoked'),
        variant: 'success',
      })
    } catch (err) {
      showToast({
        title: t('Failed to revoke access grant'),
        body:
          err.response && err.response.status === 403
            ? t(`You don't have the required permission.`)
            : err.message,
        variant: 'error',
      })
    }
  }

  return (
    <StyledCollapsibleBody>
      {accessGrants.length ? (
        <Table data-testid="service-accessgrant-list">
          <tbody>
            {accessGrants.map((accessGrant) => (
              <AccessGrantRow
                key={accessGrant.id}
                accessGrant={accessGrant}
                revokeHandler={revokeHandler}
              />
            ))}
          </tbody>
        </Table>
      ) : (
        <StyledSmall>{t('There are no organizations with access')}</StyledSmall>
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
  accessGrants: array,
  showLoadIncomingDataButton: bool,
  onClickLoadIncomingDataHandler: func,
}

// eslint-disable-next-line @typescript-eslint/no-empty-function
const noop = () => {}

CollapsibleBody.defaultProps = {
  accessGrants: [],
  showLoadIncomingDataButton: false,
  onClickLoadIncomingDataHandler: noop,
}

export default CollapsibleBody
