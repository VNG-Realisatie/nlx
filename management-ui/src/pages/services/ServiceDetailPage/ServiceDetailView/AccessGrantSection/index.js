// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { array } from 'prop-types'
import { Collapsible } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import CollapsibleBody from './CollapsibleBody'
import CollapsibleHeader from './CollapsibleHeader'

const AccessGrantSection = ({ accessGrants }) => {
  const { t } = useTranslation()

  return (
    <Collapsible
      title={<CollapsibleHeader counter={accessGrants.length} />}
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
