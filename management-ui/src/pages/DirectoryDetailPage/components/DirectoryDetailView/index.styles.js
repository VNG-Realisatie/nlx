// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Alert } from '@commonground/design-system'

import LinkButton from '../../../../components/LinkButton'

export const StyledAlert = styled(Alert)`
  margin-bottom: ${(p) => p.theme.tokens.spacing05};
`

/*
Currently not used, but will do soon
<RetryButton
  onClick={requestAccess}
  disabled={isRequestSentForThisService}
>
  <IconRedo />
  {t('Try again')}
</RetryButton>
*/
export const RetryButton = styled(LinkButton)`
  margin-top: ${(p) => p.theme.tokens.spacing05};
`

export const AccessSection = styled.section`
  display: flex;
  align-items: center;
  justify-content: space-between;
`

export const IconItem = styled.div`
  margin-right: ${(p) => p.theme.tokens.spacing03};
  color: ${(p) => p.theme.colorTextLabel};
`

export const StatusItem = styled.div`
  flex: 1 1 auto;
`

export const AccessMessage = styled.p`
  font-size: ${(p) => p.theme.tokens.fontSizeSmall};
`
