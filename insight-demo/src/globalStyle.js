import { createGlobalStyle } from 'styled-components'

const GlobalStyle = createGlobalStyle`
    html {
        font-family: ${p => p.theme.font.family.main};
        color: ${p => p.theme.color.black};
        text-rendering: optimizeLegibility;
        -webkit-font-smoothing: antialiased;
    }

    body {
        background-color: ${p => p.theme.color.white};
        margin: 0;
        overflow-wrap: break-word;
        word-wrap: break-word;
        word-break: break-word;
    }

    *,
    *:after,
    *:before {
        box-sizing: border-box;
    }

    *:active {
        transition-duration: 0s !important;
    }

    *:focus {
        outline: none;
    }
`;

export default GlobalStyle