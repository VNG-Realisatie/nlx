import React from 'react'
import { arrayOf, shape, string, instanceOf } from 'prop-types'
import LogsTable from "../LogsTable";

const rawLogs = [ {
  "source_organization": "haarlem",
  "destination_organization": "brp",
  "service_name": "basisregistratie",
  "logrecord-id": "d84ldmbot6653",
  "data": {
    "doelbinding-application-id": "Parkeervergunningapplicatie",
    "doelbinding-data-elements": "burgerservicenummer,adres",
    "doelbinding-process-id": "Aanvragen van parkeervergunning",
    "doelbinding-subject-identifier": "302641828",
    "request-path": "/ingeschreven_natuurlijke_personen"
  },
  "DataSubjects": null,
  "created": "2019-04-11T14:39:20.472528Z"
},
{
  "source_organization": "haarlem",
  "destination_organization": "brp",
  "service_name": "basisregistratie",
  "logrecord-id": "6p6029n9c9bh7",
  "data": {
    "doelbinding-application-id": "Parkeervergunningapplicatie",
    "doelbinding-data-elements": "burgerservicenummer,adres",
    "doelbinding-process-id": "Aanvragen van parkeervergunning",
    "doelbinding-subject-identifier": "302641828",
    "request-path": "/ingeschreven_natuurlijke_personen"
  },
  "DataSubjects": null,
  "created": "2019-04-11T13:07:18.98178Z"
}, {
  "source_organization": "haarlem",
  "destination_organization": "brp",
  "service_name": "basisregistratie",
  "logrecord-id": "eg38kmsv00ic0",
  "data": {
    "doelbinding-application-id": "Parkeervergunningapplicatie",
    "doelbinding-data-elements": "burgerservicenummer,adres",
    "doelbinding-process-id": "Aanvragen van parkeervergunning",
    "doelbinding-subject-identifier": "302641828",
    "request-path": "/ingeschreven_natuurlijke_personen"
  },
  "DataSubjects": null,
  "created": "2019-04-10T11:52:37.553583Z"
}, {
  "source_organization": "haarlem",
  "destination_organization": "brp",
  "service_name": "basisregistratie",
  "logrecord-id": "8popligt8cl3o",
  "data": {
    "doelbinding-application-id": "Parkeervergunningapplicatie",
    "doelbinding-data-elements": "burgerservicenummer,adres",
    "doelbinding-process-id": "Aanvragen van parkeervergunning",
    "doelbinding-subject-identifier": "302641828",
    "request-path": "/ingeschreven_natuurlijke_personen"
  },
  "DataSubjects": null,
  "created": "2019-04-10T07:20:34.575963Z"
}, {
  "source_organization": "haarlem",
  "destination_organization": "brp",
  "service_name": "basisregistratie",
  "logrecord-id": "cdd3ujeqmk1o7",
  "data": {
    "doelbinding-application-id": "Parkeervergunningapplicatie",
    "doelbinding-data-elements": "burgerservicenummer,adres",
    "doelbinding-process-id": "Aanvragen van parkeervergunning",
    "doelbinding-subject-identifier": "302641828",
    "request-path": "/ingeschreven_natuurlijke_personen"
  },
  "DataSubjects": null,
  "created": "2019-04-10T06:49:03.844605Z"
}, {
  "source_organization": "haarlem",
  "destination_organization": "brp",
  "service_name": "basisregistratie",
  "logrecord-id": "bdri3v2qqvu2a",
  "data": {
    "doelbinding-application-id": "Parkeervergunningapplicatie",
    "doelbinding-data-elements": "burgerservicenummer,adres",
    "doelbinding-process-id": "Aanvragen van parkeervergunning",
    "doelbinding-subject-identifier": "302641828",
    "request-path": "/ingeschreven_natuurlijke_personen"
  },
  "DataSubjects": null,
  "created": "2019-04-09T14:12:59.718246Z"
}, {
  "source_organization": "haarlem",
  "destination_organization": "brp",
  "service_name": "basisregistratie",
  "logrecord-id": "ecd50p1lcmvdc",
  "data": {
    "doelbinding-application-id": "Parkeervergunningapplicatie",
    "doelbinding-data-elements": "burgerservicenummer,adres",
    "doelbinding-process-id": "Aanvragen van parkeervergunning",
    "doelbinding-subject-identifier": "302641828",
    "request-path": "/ingeschreven_natuurlijke_personen"
  },
  "DataSubjects": null,
  "created": "2019-04-09T09:49:33.268826Z"
}, {
  "source_organization": "haarlem",
  "destination_organization": "brp",
  "service_name": "basisregistratie",
  "logrecord-id": "6cqegqnckrka6",
  "data": {
    "doelbinding-application-id": "Parkeervergunningapplicatie",
    "doelbinding-data-elements": "burgerservicenummer,adres",
    "doelbinding-process-id": "Aanvragen van parkeervergunning",
    "doelbinding-subject-identifier": "302641828",
    "request-path": "/ingeschreven_natuurlijke_personen"
  },
  "DataSubjects": null,
  "created": "2019-04-05T08:56:52.523184Z"
}, {
  "source_organization": "haarlem",
  "destination_organization": "brp",
  "service_name": "basisregistratie",
  "logrecord-id": "9v2gt7n2b7ive",
  "data": {
    "doelbinding-application-id": "Parkeervergunningapplicatie",
    "doelbinding-data-elements": "burgerservicenummer,adres",
    "doelbinding-process-id": "Aanvragen van parkeervergunning",
    "doelbinding-subject-identifier": "302641828",
    "request-path": "/ingeschreven_natuurlijke_personen"
  },
  "DataSubjects": null,
  "created": "2019-04-04T14:32:19.440555Z"
}, {
  "source_organization": "haarlem",
  "destination_organization": "brp",
  "service_name": "basisregistratie",
  "logrecord-id": "akt7dhi1fi41g",
  "data": {
    "doelbinding-application-id": "Parkeervergunningapplicatie",
    "doelbinding-data-elements": "burgerservicenummer,adres",
    "doelbinding-process-id": "Aanvragen van parkeervergunning",
    "doelbinding-subject-identifier": "302641828",
    "request-path": "/ingeschreven_natuurlijke_personen"
  },
  "DataSubjects": null,
  "created": "2019-04-04T10:55:15.617865Z"
}
]

const LogsPage = ({ logs }) =>
  <div>
    <LogsTable logs={logs} />
  </div>

LogsPage.propTypes = {
  logs: arrayOf(shape({
    subjects: arrayOf(string),
    requestedBy: string,
    requestedAt: string,
    reason: string,
    date: instanceOf(Date)
  }))
}

const mapRawLogsToTableFormat = rawLogs =>
  rawLogs.map(rawLog => ({
    subjects: rawLog.data['doelbinding-data-elements'].split(','),
    requestedBy: rawLog['source_organization'],
    requestedAt: rawLog['destination_organization'],
    reason: rawLog.data['doelbinding-process-id'],
    date: new Date(rawLog['created'])
  }))

LogsPage.defaultProps = {
  logs: mapRawLogsToTableFormat(rawLogs)
}

export default LogsPage
