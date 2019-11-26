/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

import React from 'react'
import Link from '@docusaurus/Link'
import useDocusaurusContext from '@docusaurus/useDocusaurusContext'
import useBaseUrl from '@docusaurus/useBaseUrl'
import {Redirect} from 'react-router-dom'


function Home() {
    const context = useDocusaurusContext()
    const {siteConfig = {}} = context
    const url = useBaseUrl(siteConfig.customFields.startUrl)
    return (<>
            <Redirect to={url}/>
            <p>
                If you are not redirected automatically, follow this <Link to={url}>link</Link>.
            </p>
        </>
    )
}

export default Home
