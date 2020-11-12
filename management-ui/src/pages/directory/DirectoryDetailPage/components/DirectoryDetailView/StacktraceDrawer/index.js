// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, string } from 'prop-types'
import { Drawer, StackedDrawer } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import { StyledPre } from './index.styles'

const StacktraceDrawer = ({ stacktrace, ...props }) => {
  const { t } = useTranslation()

  return (
    <StackedDrawer {...props} data-testid="stacktrace">
      <Drawer.Header as="header" title={t('Stacktrace')} />
      <Drawer.Content>
        <StyledPre data-testid="stacktrace-content">
          {stacktrace.map((line, i) => (
            <code key={i}>
              {line}
              <br />
            </code>
          ))}
        </StyledPre>
      </Drawer.Content>
    </StackedDrawer>
  )
}

StacktraceDrawer.propTypes = {
  id: string.isRequired,
  parentId: string.isRequired,
  stacktrace: arrayOf(string),
}

StacktraceDrawer.defaultProps = {}

export default StacktraceDrawer
