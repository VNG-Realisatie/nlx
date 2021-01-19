// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { number } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { IconKey } from '../../../../../../icons'
import Amount from '../../../../../../components/Amount'
import { DetailHeading } from '../../../../../../components/DetailView'

const CollapsibleHeader = ({ counter }) => {
  const { t } = useTranslation()
  return (
    <DetailHeading>
      <IconKey />
      {t('Access requests')}
      <Amount value={counter} isAccented={counter > 0} data-testid="amount" />
    </DetailHeading>
  )
}

CollapsibleHeader.propTypes = {
  counter: number,
}

CollapsibleHeader.defaultProps = {
  counter: 0,
}

export default CollapsibleHeader
