// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { Icon } from '@commonground/design-system'

import { ReactComponent as IconChevronRight } from './chevron-right.svg'
import { ReactComponent as IconExternalLink } from './external-link.svg'
import { ReactComponent as IconHome } from './home.svg'
import { ReactComponent as IconMail } from './mail.svg'
import { ReactComponent as IconStateDegraded } from './state-degraded.svg'
import { ReactComponent as IconStateDown } from './state-down.svg'
import { ReactComponent as IconStateUnknown } from './state-unknown.svg'
import { ReactComponent as IconStateUp } from './state-up.svg'
import { ReactComponent as MoneyEuroCircleLine } from './money-euro-circle-line.svg'

export {
  IconChevronRight,
  IconExternalLink,
  IconHome,
  IconMail,
  IconStateDegraded,
  IconStateDown,
  IconStateUnknown,
  IconStateUp,
}

export const IconMoneyEuroCircleLine = (props) => (
  <Icon as={MoneyEuroCircleLine} {...props} />
)
