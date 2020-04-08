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
  
  button,
  a {
    border: 2px solid transparent;
    
    &:focus {
      outline: none;
      border-color: ${(p) => p.theme.colorBorderDropdownFocus};
    }
  }  
`
