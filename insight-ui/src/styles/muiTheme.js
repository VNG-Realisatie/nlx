import { createMuiTheme } from '@material-ui/core/styles'
import 'typeface-source-sans-pro'
import 'typeface-muli'

export const muiTheme = createMuiTheme({
    typography: {
        useNextVariants: true,
        fontFamily: "'Source Sans Pro', 'Helvetica', 'Arial', sans-serif",
        fontSize: 16,
        fontWeightLight: 300,
        fontWeightRegular: 400,
        fontWeightMedium: 600,
        lineHeight: 1.5,
        h4: {
            fontFamily: 'Muli, sans-serif',
            letterSpacing: 4,
            textTransform: 'uppercase',
            fontWeight: 700,
        },
        // previously called headline
        h5: {
            fontFamily: "'Muli', 'Helvetica', 'Arial', sans-serif",
            fontWeight: 600,
            letterSpacing: -1,
        },
        // previously called title
        h6: {
            fontFamily: "'Muli', 'Helvetica', 'Arial', sans-serif",
            fontWeight: 600,
            letterSpacing: -1,
        },
    },
    spacing: {
        unit: 16,
    },
    palette: {
        common: {
            black: '#000',
            white: '#fff',
        },
        primary: {
            main: '#FEBF24',
        },
        secondary: {
            main: '#3D83FA',
        },
        text: {
            primary: 'rgba(0, 0, 0, 0.75)',
            secondary: 'rgba(0, 0, 0, 0.4)',
            disabled: 'rgba(0, 0, 0, 0.25)',
            hint: 'rgba(0, 0, 0, 0.4)',
            divider: 'rgba(0, 0, 0, 0.1)',
        },
        background: {
            paper: '#fff',
            default: '#fff',
        },
        action: {
            active: 'rgba(0, 0, 0, 0.5)',
            hover: 'rgba(0, 0, 0, 0.03)',
            hoverOpacity: 0.08,
            selected: 'rgba(0, 0, 0, 0.5)',
            disabled: 'rgba(0, 0, 0, 0.25)',
            disabledBackground: 'rgba(0, 0, 0, 0.1)',
        },
    },
    shadows: [
        'none',
        '0px 1px 3px 0px rgba(0, 0, 0, 0.05),0px 1px 1px 0px rgba(0, 0, 0, 0.02),0px 2px 1px -1px rgba(0, 0, 0, 0.03)',
        '0px 1px 5px 0px rgba(0, 0, 0, 0.05),0px 2px 2px 0px rgba(0, 0, 0, 0.02),0px 3px 1px -2px rgba(0, 0, 0, 0.03)',
        '0px 1px 8px 0px rgba(0, 0, 0, 0.05),0px 3px 4px 0px rgba(0, 0, 0, 0.02),0px 3px 3px -2px rgba(0, 0, 0, 0.03)',
        '0px 2px 4px -1px rgba(0, 0, 0, 0.05),0px 4px 5px 0px rgba(0, 0, 0, 0.02),0px 1px 10px 0px rgba(0, 0, 0, 0.03)',
        '0px 3px 5px -1px rgba(0, 0, 0, 0.05),0px 5px 8px 0px rgba(0, 0, 0, 0.02),0px 1px 14px 0px rgba(0, 0, 0, 0.03)',
        '0px 3px 5px -1px rgba(0, 0, 0, 0.05),0px 6px 10px 0px rgba(0, 0, 0, 0.02),0px 1px 18px 0px rgba(0, 0, 0, 0.03)',
        '0px 4px 5px -2px rgba(0, 0, 0, 0.05),0px 7px 10px 1px rgba(0, 0, 0, 0.02),0px 2px 16px 1px rgba(0, 0, 0, 0.03)',
        '0px 5px 5px -3px rgba(0, 0, 0, 0.05),0px 8px 10px 1px rgba(0, 0, 0, 0.02),0px 3px 14px 2px rgba(0, 0, 0, 0.03)',
        '0px 5px 6px -3px rgba(0, 0, 0, 0.05),0px 9px 12px 1px rgba(0, 0, 0, 0.02),0px 3px 16px 2px rgba(0, 0, 0, 0.03)',
        '0px 6px 6px -3px rgba(0, 0, 0, 0.05),0px 10px 14px 1px rgba(0, 0, 0, 0.02),0px 4px 18px 3px rgba(0, 0, 0, 0.03)',
        '0px 6px 7px -4px rgba(0, 0, 0, 0.05),0px 11px 15px 1px rgba(0, 0, 0, 0.02),0px 4px 20px 3px rgba(0, 0, 0, 0.03)',
        '0px 7px 8px -4px rgba(0, 0, 0, 0.05),0px 12px 17px 2px rgba(0, 0, 0, 0.02),0px 5px 22px 4px rgba(0, 0, 0, 0.03)',
        '0px 7px 8px -4px rgba(0, 0, 0, 0.05),0px 13px 19px 2px rgba(0, 0, 0, 0.02),0px 5px 24px 4px rgba(0, 0, 0, 0.03)',
        '0px 7px 9px -4px rgba(0, 0, 0, 0.05),0px 14px 21px 2px rgba(0, 0, 0, 0.02),0px 5px 26px 4px rgba(0, 0, 0, 0.03)',
        '0px 8px 9px -5px rgba(0, 0, 0, 0.05),0px 15px 22px 2px rgba(0, 0, 0, 0.02),0px 6px 28px 5px rgba(0, 0, 0, 0.03)',
        '0px 8px 10px -5px rgba(0, 0, 0, 0.05),0px 16px 24px 2px rgba(0, 0, 0, 0.02),0px 6px 30px 5px rgba(0, 0, 0, 0.03)',
        '0px 8px 11px -5px rgba(0, 0, 0, 0.05),0px 17px 26px 2px rgba(0, 0, 0, 0.02),0px 6px 32px 5px rgba(0, 0, 0, 0.03)',
        '0px 9px 11px -5px rgba(0, 0, 0, 0.05),0px 18px 28px 2px rgba(0, 0, 0, 0.02),0px 7px 34px 6px rgba(0, 0, 0, 0.03)',
        '0px 9px 12px -6px rgba(0, 0, 0, 0.05),0px 19px 29px 2px rgba(0, 0, 0, 0.02),0px 7px 36px 6px rgba(0, 0, 0, 0.03)',
        '0px 10px 13px -6px rgba(0, 0, 0, 0.05),0px 20px 31px 3px rgba(0, 0, 0, 0.02),0px 8px 38px 7px rgba(0, 0, 0, 0.03)',
        '0px 10px 13px -6px rgba(0, 0, 0, 0.05),0px 21px 33px 3px rgba(0, 0, 0, 0.02),0px 8px 40px 7px rgba(0, 0, 0, 0.03)',
        '0px 10px 14px -6px rgba(0, 0, 0, 0.05),0px 22px 35px 3px rgba(0, 0, 0, 0.02),0px 8px 42px 7px rgba(0, 0, 0, 0.03)',
        '0px 11px 14px -7px rgba(0, 0, 0, 0.05),0px 23px 36px 3px rgba(0, 0, 0, 0.02),0px 9px 44px 8px rgba(0, 0, 0, 0.03)',
        '0px 11px 15px -7px rgba(0, 0, 0, 0.05),0px 24px 38px 3px rgba(0, 0, 0, 0.02),0px 9px 46px 8px rgba(0, 0, 0, 0.03)',
    ],
    overrides: {
        MuiAppBar: {
            colorDefault: {
                backgroundColor: '#ffffff',
                color: '#000000',
            },
        },
        MuiDrawer: {
            paperAnchorDockedLeft: {
                width: 248,
                borderRight: 'none',
            },
            paper: {
                backgroundColor: '#FAFAFA',
            },
        },
        MuiListSubheader: {
            root: {
                outline: 'none',
                pointer: 'default',
            },
        },
        MuiMenuItem: {
            root: {
                fontSize: '1rem',
            },
            selected: {
                fontWeight: 700,
            },
        },
        MuiListItem: {
            root: {
                paddingTop: 8,
                paddingBottom: 8,
                '&$selected': {
                    backgroundColor: 'transparent',
                    color: '#3D83FA',
                },
            },
        },
        MuiListItemText: {
            root: {
                paddingRight: 0,
                '& span': {
                    fontSize: '1rem',
                    overflow: 'hidden',
                    textOverflow: 'ellipsis',
                },
            },
        },
        MuiButton: {
            root: {
                textTransform: 'none',
            },
            contained: {
                boxShadow: 'none',
                '&:active': {
                    boxShadow: 'none',
                },
            },
            containedPrimary: {
                boxShadow: '0 3px 0 #eea901',
                '&:active': {
                    boxShadow: '0 3px 0 #eea901',
                },
            },
        },
        MuiIconButton: {
            root: {
                width: 40,
                height: 40,
            },
        },
        MuiTypography: {
            gutterBottom: {
                marginBottom: '1.5rem',
            },
            body1: {
                fontSize: '1rem',
                lineHeight: 1.5,
            },
            body2: {
                fontSize: '1rem',
                lineHeight: 1.5,
                fontWeight: 600,
            },
            caption: {
                fontSize: '0.875rem',
            },
        },
        MuiTableBody: {
            root: {
                backgroundColor: '#ffffff',
            },
        },
        MuiTableRow: {
            root: {
                borderLeft: '1px solid #eeeeee',
                borderRight: '1px solid #eeeeee',
            },
            head: {
                height: '48px',
                borderColor: 'transparent',
            },
        },
        MuiTableCell: {
            root: {
                padding: '4px 24px',
                borderBottom: '1px solid #eeeeee',
                '&:first-child': {
                    textAlign: 'center',
                },
            },
            body: {
                fontSize: '0.875rem',
            },
            head: {
                paddingTop: 0,
                paddingBottom: 0,
                fontSize: '0.875rem',
                fontWeight: '400',
            },
        },
        MuiPopover: {
            paper: {
                transform: 'none !important',
            },
        },
        MuiTableSortLabel: {
            root: {
                height: '48px',
            },
            iconDirectionAsc: {
                transform: 'rotate(0deg)',
            },
            iconDirectionDesc: {
                transform: 'rotate(180deg)',
            },
        },
        MuiTablePagination: {
            root: {
                fontSize: '.875rem',
            },
            caption: {
                fontSize: '.875rem',
            },
            selectIcon: {
                fontSize: '17px',
                top: 5,
                right: 2,
            },
        },
        MuiSelect: {
            selectMenu: {
                minHeight: 'auto',
                paddingBottom: 6,
            },
        },
        MuiBackdrop: {
            root: {
                cursor: 'pointer',
            },
        },
    },
})

export const globalStyles = (theme) => ({
    '@global': {
        body: {
            fontSize: theme.typography.fontSize,
            lineHeight: theme.typography.lineHeight,
            fontFamily: theme.typography.fontFamily,
        },
        a: {
            textDecoration: 'none',
            color: '#3D83FA',
            transition: 'color .15s ease',
            '&:hover': {
                color: '#81aefc',
            },
        },
        p: {
            marginTop: 0,
        },
    },
})

export default muiTheme
