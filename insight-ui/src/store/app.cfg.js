// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

/**
 * Inital app configuration
 * Values are imported into redux store
 */
const cfg = {
    /**
     * For testing in local env with IRMA app see README file
     * for more info (line 57 onwards). Here you need to
     * uncomment property "localIp" and provide local ip
     * where insight-ui app is running, incl. port 3333
     * Note: provided ip cannot be 'localhost:3333'     *
     * IMPORTANT!
     * before commiting you need to comment/disable localIp prop
     * otherwise CI/CD tests will fail
     */
    // localIp: 'http://192.168.7.77:3333',
    //
    // app-loader
    loader: {
        show: true,
    },
    //
    // internationalization
    i18n: {
        defaultLang: 'en',
        // localStorage key
        lsKey: 'nlx.app.lang',
        // list of languages
        options: [
            {
                key: 'en',
                // i18next ns=namespace
                ns: 'core',
                label: 'English',
                data: 'locale/en/us.json',
                icon: 'img/us.svg',
            },
            {
                key: 'nl',
                // i18next ns=namespace
                ns: 'core',
                label: 'Nederlands',
                data: 'locale/nl/nl.json',
                icon: 'img/nl.svg',
            },
        ],
        // current language info goes here
        // see languageReducer for implementation
        lang: {
            key: null,
            label: null,
            data: null,
        },
    },
    //
    // list of all organizations
    organizations: {
        api: '/api/directory/list-organizations',
        list: [],
        error: null,
    },
    //
    // currently loaded organization
    organization: {
        info: {
            name: null,
            // eslint-disable-next-line
            insight_irma_endpoint: null,
            // eslint-disable-next-line
            insight_log_endpoint: null,
        },
        irma: {
            name: null,
            dataSubjects: null,
            qrCode: null,
            statusUrl: null,
            proofUrl: null,
            firstJWT: null,
            jwt: null,
            error: null,
            inProgress: false,
        },
        logs: {
            name: null,
            api: null,
            error: null,
            items: [],
            // column definitions for logs table
            colDef: [
                {
                    id: 'date',
                    label: 'Datum',
                    width: 100,
                    src: 'created',
                    type: 'date',
                    disablePadding: true,
                },
                {
                    id: 'time',
                    label: 'Tijd',
                    src: 'created',
                    type: 'time',
                    disablePadding: false,
                },
                {
                    id: 'source',
                    label: 'Opgevraagd door',
                    src: 'source_organization',
                    type: 'string',
                    disablePadding: false,
                },
                {
                    id: 'destination',
                    label: 'Opgevraagd bij',
                    src: 'destination_organization',
                    type: 'string',
                    disablePadding: false,
                },
                {
                    id: 'reden',
                    label: 'Process',
                    src: 'data.doelbinding-process-id',
                    type: 'string',
                    disablePadding: false,
                },
            ],
            pageDef: {
                page: 0,
                rowsPerPage: 10,
                rowCount: 0,
                rowsPerPageOptions: [5, 10, 25, 50],
            },
        },
    },
}

export default cfg
