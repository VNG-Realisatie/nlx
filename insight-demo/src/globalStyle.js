import { createGlobalStyle } from 'styled-components'

/*
    Dusan, 2018-12-07
    GoogleFonts css file uses woff2 files instead of ttf. The difference in the level
    of support is quite small but file size difference is signficant. I advice woff2.
*/

import sourceSansProRegularWoff2 from './fonts/sourcesanspro/SourceSansPro-Regular.woff2';
import sourceSansProSemiBoldWoff2 from './fonts/sourcesanspro/SourceSansPro-SemiBold.woff2';
import sourceSansProBoldWoff2 from './fonts/sourcesanspro/SourceSansPro-Bold.woff2';

import sourceSansProRegularWoff from './fonts/sourcesanspro/SourceSansPro-Regular.woff';
import sourceSansProSemiBoldWoff from './fonts/sourcesanspro/SourceSansPro-SemiBold.woff';
import sourceSansProBoldWoff from './fonts/sourcesanspro/SourceSansPro-Bold.woff';

const GlobalStyle = createGlobalStyle`
    @font-face {
        font-family: '${p => p.theme.font.family.main}';
        src:
            url(${sourceSansProRegularWoff2}) format('woff2'),
            url(${sourceSansProRegularWoff}) format('woff');
        font-weight: ${p => p.theme.font.weight.normal};
        font-style: normal;
    }

    @font-face {
        font-family: '${p => p.theme.font.family.main}';
        src:
            url(${sourceSansProSemiBoldWoff2}) format('woff2'),
            url(${sourceSansProSemiBoldWoff}) format('woff');
        font-weight: ${p => p.theme.font.weight.semibold};
        font-style: normal;
    }

    @font-face {
        font-family: '${p => p.theme.font.family.main}';
        src:
            url(${sourceSansProBoldWoff2}) format('woff2'),
            url(${sourceSansProBoldWoff}) format('woff');
        font-weight: ${p => p.theme.font.weight.bold};
        font-style: normal;
    }

    body {
        font-family: '${p => p.theme.font.family.main}';
        color: ${p => p.theme.color.black};
        background-color: ${p => p.theme.color.white};
        margin: 0;
        overflow-wrap: break-word;
        word-wrap: break-word;
        word-break: break-word;
        text-rendering: optimizeLegibility;
        -webkit-font-smoothing: antialiased;
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