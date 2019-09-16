// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React, { useState } from 'react'
import { func } from 'prop-types'
import { Button } from '../../components/Form'
import { StyledConfirm } from './index.styles'

const ConfirmButton = ({ onConfirm }) => {
    const [confirm, setConfirm] = useState(false)

    return (
        <div>
            {confirm ? (
                <StyledConfirm>
                    <div>Are you sure?</div>
                    <Button
                        alert
                        type="button"
                        onClick={onConfirm}
                        data-test="confirm-button"
                    >
                        Confirm delete
                    </Button>
                    <Button
                        secondary
                        type="button"
                        onClick={() => setConfirm(false)}
                        data-test="cancel-button"
                    >
                        Cancel
                    </Button>
                </StyledConfirm>
            ) : (
                <Button
                    secondary
                    alert
                    type="button"
                    onClick={() => setConfirm(true)}
                    data-test="delete-button"
                >
                    Delete
                </Button>
            )}
        </div>
    )
}

ConfirmButton.propTypes = {
    onConfirm: func.isRequired,
}

export default ConfirmButton
