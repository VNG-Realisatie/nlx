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
`
