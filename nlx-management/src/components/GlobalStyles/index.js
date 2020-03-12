import { createGlobalStyle } from 'styled-components'

export default createGlobalStyle`
  html,
  body,
  #root {
    height: 100%;
  }
  
  #root {
    display: flex;
    flex-direction: column;
  }
`
