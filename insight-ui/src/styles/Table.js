// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

export const tableStyles = (theme) => ({
    tableWrapper: {
        overflowX: 'auto',
    },
    firstTableCell: {
        width: 60,
        textAlign: 'center',
    },
    iconTableCell: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
    },
    calendarIcon: {
        fontSize: 14,
        marginBottom: -2,
        marginRight: 3,
    },
    clockIcon: {
        fontSize: 15,
        marginBottom: -3,
        marginRight: 2,
    },
    MuiTableCell: {
        body: {
            wordBreak: 'break-word',
        },
    },
})

export default tableStyles
