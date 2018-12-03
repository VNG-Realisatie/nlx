import React from 'react'
import { configure, addDecorator } from '@storybook/react';
import { createGlobalStyle, ThemeProvider } from 'styled-components'
import theme from './theme'

import { setDefaults } from '@storybook/addon-info'
import { withOptions } from '@storybook/addon-options'

// addon-info
setDefaults({
    header: false, // Toggles display of header with component name and description
    inline: true, // Displays info inline vs click button to view
    source: false, // Displays the source of story Component
    styles: {
        infoBody: {
            border: 'none',
            borderTop: '1px solid rgb(238, 238, 238)',
            padding: '40px',
            backgroundColor: '#fff',
            marginTop: '0',
            marginBottom: '20px',
            boxShadow: 'none'
        },
        infoStory: {
            backgroundColor: '#fff',
            padding: '40px 40px 30px',
        },
        infoContent: {
            marginBottom: '30px'
        },
        propTableHead: {
            display: 'none',
        },
        source: {
            h1: {
                display: 'none',
            },
        },
    },
});


// Option defaults:
addDecorator(
  withOptions({
    showAddonPanel: false,
    sidebarAnimations: true,
  })
);

function loadStories() {
    require('./stories/atoms.stories.js');
    require('./stories/modules.stories.js');
}

// Global styles but theme- and update-able!
const GlobalStyle = createGlobalStyle`
    html {
        font-family: ${p => p.theme.font.family.body};
        color: ${p => p.theme.color.black};
        text-rendering: optimizeLegibility;
        -webkit-font-smoothing: antialiased;
    }

    body {
        background-color: ${p => p.theme.color.grey[10]};
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

    input, textarea, select, button {
        font: inherit;
    }
`;

addDecorator((story) => (
    <ThemeProvider theme={theme}>
        <React.Fragment>
            <GlobalStyle />
            {story()}
        </React.Fragment>
    </ThemeProvider>
))

configure(loadStories, module)