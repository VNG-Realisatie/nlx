import { createGlobalStyle } from 'styled-components'

/*
    Dusan, 2018-12-07
    GoogleFonts css file uses woff2 files instead of ttf. The difference in the level
    of support is quite small but file size difference is signficant. I advice woff2.
*/
import sourceSansProRegular from './fonts/sourcesanspro/6xK3dSBYKcSV-LCoeQqfX1RYOo3qOK7l.woff2';
import sourceSansProSemiBold from './fonts/sourcesanspro/6xKydSBYKcSV-LCoeQqfX1RYOo3i54rwlxdu.woff2';
import sourceSansProBold from './fonts/sourcesanspro/6xKydSBYKcSV-LCoeQqfX1RYOo3ig4vwlxdu.woff2';


const GlobalStyle = createGlobalStyle`
    @font-face {
        font-family: '${p => p.theme.font.family.main}';
        src: url(${sourceSansProRegular}) format('woff2');
        font-weight: ${p => p.theme.font.weight.normal};
        font-style: normal;
    }

    @font-face {
        font-family: '${p => p.theme.font.family.main}';
        src: url(${sourceSansProSemiBold}) format('woff2');
        font-weight: ${p => p.theme.font.weight.semibold};
        font-style: normal;
    }

    @font-face {
        font-family: '${p => p.theme.font.family.main}';
        src: url(${sourceSansProBold}) format('woff2');
        font-weight: ${p => p.theme.font.weight.bold};
        font-style: normal;
    }

    html {
        font-family: '${p => p.theme.font.family.main}';
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