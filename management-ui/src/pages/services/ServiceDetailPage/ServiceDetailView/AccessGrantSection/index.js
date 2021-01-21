// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { array } from 'prop-types'
import { Collapsible } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { DetailHeading } from '../../../../../components/DetailView'
import Amount from '../../../../../components/Amount'
import { IconCheckboxMultiple } from '../../../../../icons'
import CollapsibleBody from './CollapsibleBody'

const AccessGrantSection = ({ accessGrants }) => {
  const { t } = useTranslation()

  return (
    <Collapsible
      title={
        <DetailHeading data-testid="service-accessgrants">
          <IconCheckboxMultiple />
          {t('Organizations with access')}
          <Amount value={accessGrants.length} />
        </DetailHeading>
      }
      ariaLabel={t('Organizations with access')}
    >
      <CollapsibleBody accessGrants={accessGrants} />
    </Collapsible>
  )
}

AccessGrantSection.propTypes = {
  accessGrants: array,
}
AccessGrantSection.defaultProps = {
  accessGrants: [],
}

export default AccessGrantSection
