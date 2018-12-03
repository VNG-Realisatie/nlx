import { lighten, setLightness, darken } from 'polished'

export default {
    font: {
        family: {
            system: '-apple-system,BlinkMacSystemFont,"Segoe UI",Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Segoe UI Symbol"',
            body: '"Source Sans Pro", sans-serif',
        },
        size: {
            large: '18px',
            normal: '16px',
            small: '14px',
            tiny: '12px',
            title: {
                large: '36px',
                normal: '24px',
                small: '20px',
            }
        },
        lineHeight: {
            normal: '24px',
            small: '20px',
            tiny: '20px',
            title: {
                large: '44px',
                normal: '32px',
                small: '28px',
            }
        },
        weight: {
            normal: '400',
            semibold: '600',
            bold: '700'
        },
        offset: {
            top: 0,
            bottom: '2px',
        },
    },
    color: {
        primary: {
            main: '#FEBF24',
            dark: darken(.03, '#FEBF24'),
            light: lighten(.06, '#FEBF24'),
            lighter: setLightness(.9, '#FEBF24'),
            lightest: setLightness(.95, '#FEBF24'),
        },
        secondary: {
            main: '#1B6CF8', // maybe 5656FC
            light: lighten(.06, '#1B6CF8'),
        },
        white: '#FFFFFF',
        black: '#212121',
        grey: {
            10: '#F9F9F9',  // body background
            20: '#F5F5F5',  // disabled background
            30: '#EAEAEA',  // input border
            40: '#DADADA',  // input border focus
            50: '#B4B4B4',  // label, helper
            60: '#757575',  // readable grey
        },
        alert: 'rgb(249, 71, 71)',
        accept: 'rgb(84, 194, 119)',
        hover: 'rgba(0,0,0,.025)',
        active: 'rgba(0,0,0,.04)',
    },
    size: {
        small: '32px',
        normal: '40px',
        large: '48px',
        header: '48px',
    },
    offset: {
        button: '2px',
    },
    radius: {
        small: '5px',
    },
    transition: {
        fast: '0.15s ease',
        normal: '0.25s ease',
        materialNormal: '0.25s cubic-bezier(0.4, 0, 0.2, 1)',
        materialSlow: '0.35s cubic-bezier(0.4, 0, 0.2, 1)',
    },
    space: [4, 8, 16, 24, 32],
    containerWidth: '1140px',
    breakpoints: [
        '576px', // Small devices (landscape phones, 576px and up)
        '768px', // Medium devices (tablets, 768px and up)
        '992px', // Large devices (desktops, 992px and up)
        '1200px', // Extra large devices (large desktops, 1200px and up)
    ],
}