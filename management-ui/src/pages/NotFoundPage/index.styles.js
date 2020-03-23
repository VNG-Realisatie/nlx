import styled from 'styled-components'

import iconAlertCircle from './iconAlertCircle.svg'

export const NotFoundContainer = styled.section`
  position: relative;
  margin-top: 100px;
  margin-left: 242px;
  max-width: 420px;

  &::before {
    content: '';
    position: absolute;
    top: ${(p) => p.theme.tokens.spacing04};
    left: -${(p) => p.theme.tokens.spacing11};
    z-index: -1;
    width: 51px;
    height: 51px;
    background: url(${iconAlertCircle}) no-repeat;
  }
`
