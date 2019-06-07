// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { useState, useEffect } from 'react'
import ConsoleLog from '../../components/ConsoleLog/ConsoleLog'
 
function Version() {
    const [data, setData] = useState({
        tag: undefined
    })
    
    useEffect(
        () => {
        const load = async () => {
            const result = await fetch('/version.json')
            const data = await result.json()
            setData(data)
        }
        load()
        }, 
        [], // prevent inifinite rerenders
    );
    
    return (
        <>{ 
            data.tag && <ConsoleLog>{data.tag}</ConsoleLog>
        } </>
    )
}

export default Version
