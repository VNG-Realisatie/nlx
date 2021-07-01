// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { createGlobalStyle } from 'styled-components'

export default createGlobalStyle`
  html,
  body,
  #root {
    height: 100vh;
  }

  #root {
    display: flex;
    flex-direction: column;
  }

  *[tabindex], a, button {
    :focus {
      outline: 2px solid ${(p) => p.theme.colorFocus};
    }
  }
`
