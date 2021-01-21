// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { node } from 'prop-types'
import { Icon } from '@commonground/design-system'
import { IconArrowDownLine } from '../../../../../icons'
import Button from './index.styles'

const UpdateUiButton = ({ children, ...props }) => (
  <Button {...props}>
    <Icon as={IconArrowDownLine} inline />
    {children}
  </Button>
)

UpdateUiButton.propTypes = {
  children: node,
}

export default UpdateUiButton
