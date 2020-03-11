import styled from 'styled-components'
import { NLXLogo } from '@commonground/design-system'
import nlxPattern from './nlx-pattern.svg'

export const StyledNLXLogo = styled(NLXLogo)`
  width: 100px;
`

export const StyledContainer = styled.div`
  display: flex;
  align-items: center;
  height: 100%;
`

export const StyledSidebar = styled.div`
  background: #313131;
  background-image: url(${nlxPattern});
  flex: 0 0 532px;
  text-align: center;
  font-size: 14pt;
  padding-top: 175px;
  height: 100%;
`

export const StyledContent = styled.div`
  flex: 1;
  padding: 0 180px;
`
