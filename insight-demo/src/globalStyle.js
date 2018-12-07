import { createGlobalStyle } from 'styled-components'
import sourceSansProRegular from './fonts/sourcesanspro/SourceSansPro-Regular.ttf';
import sourceSansProSemiBold from './fonts/sourcesanspro/SourceSansPro-SemiBold.ttf';
import sourceSansProBold from './fonts/sourcesanspro/SourceSansPro-Bold.ttf';


const GlobalStyle = createGlobalStyle`
    @font-face {
        font-family: ${p => p.theme.font.family.main};
        src: url(${sourceSansProRegular}) format('truetype');
        font-weight: ${p => p.theme.font.weight.normal};
        font-style: normal;
    }

    @font-face {
        font-family: ${p => p.theme.font.family.main};
        src: url(${sourceSansProSemiBold}) format('truetype');
        font-weight: ${p => p.theme.font.weight.semibold};
        font-style: normal;
    }

    @font-face {
        font-family: ${p => p.theme.font.family.main};
        src: url(${sourceSansProBold}) format('truetype');
        font-weight: ${p => p.theme.font.weight.bold};
        font-style: normal;
    }

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
        flex-shrink: 0;
    }

    *:active {
        transition-duration: 0s !important;
    }

    *:focus {
        outline: none;
    }
`;

export default GlobalStyle