const modalStyles = (theme) => ({
    paper: {
        position: 'fixed',
        top: '50%',
        left: '50%',
        transform: 'translate(-50%, -50%)',
        width: 540,
        maxWidth: '100%',
        padding: `${theme.spacing.unit * 2.25}px ${theme.spacing.unit * 2}px`,
        overflowY: 'auto',
        '@media (max-width: 767px)': {
            top: 0,
            left: 0,
            right: 0,
            bottom: 0,
            transform: 'translate(0, 0)',
            width: 'auto',
            borderRadius: 0,
        },
    },
    closeButton: {
        position: 'absolute',
        top: 12,
        right: 15,
    },
})

export default modalStyles
