// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'

const Button = styled.button`
  background-color: ${(p) => p.theme.tokens.colorBrand4};
  color: ${(p) => p.theme.tokens.colorTextInverse};
  border-radius: 1rem;
  height: 2rem;
  padding: 0 1rem;
  cursor: pointer;
  display: block;
  position: relative;
  margin: -2rem auto 0 auto;
  top: 6px;
`

export default Button
