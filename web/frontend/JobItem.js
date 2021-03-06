// Copyright 2018 Delivery Hero Germany
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
//     Unless required by applicable law or agreed to in writing, software
//     distributed under the License is distributed on an "AS IS" BASIS,
//     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//     See the License for the specific language governing permissions and
//     limitations under the License.

import React, { Component } from 'react';
import './App.css';
import config from "./config";
import { Col } from "react-bootstrap";

export default class JobItem extends Component {
    constructor(props) {
        super(props);

        this.state = {
            job: props.job,
            username: null
        }
    }

    async fetchUsername() {
        try {
            const response = await fetch('/api/user/' + this.state.job.author, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + localStorage.getItem(config.OBI_TOKEN_KEY)
                }
            });
            let user = await response.json();
             this.setState({
                username: user.email
            })
        }
        catch (err) {
            this.setState({
                username: null
            })
        }
    }

    async UNSAFE_componentWillMount() {
        await this.fetchUsername()
    }

    render() {
        return (
            <div className="JobItem">
                <Col md={3} sm={12}>
                    <span className="JobItem-id"><b>ID</b>: {this.state.job.id}</span>
                </Col>
                <Col md={3} sm={6}>
                    <span className="JobItem-status"><b>Status</b>: {this.state.job.status}</span>
                </Col>
                <Col md={6} sm={6}>
                    <span className="JobItem-user"><b>Submitted by</b>: {this.state.username}</span>
                </Col>
                <Col md={6} sm={6}>
                    <span className="JobItem-executablepath"><b>Executable</b>: {this.state.job.executablepath}</span>
                </Col>
                <Col md={6} sm={6}>
                    <span className="JobItem-executablepath"><b>Arguments</b>: {this.state.job.arguments}</span>
                </Col>
                <Col md={12} sm={12}>
                    <span className="JobItem-link">
                        <a href={"https://console.cloud.google.com/dataproc/jobs/" +
                                    this.state.job.platformdependentid + "?region=global"}
                           target="_blank" rel="noopener noreferrer">
                            Job Logs</a></span>
                </Col>
                <Col md={12} sm={12}>
                    <hr/>
                </Col>
            </div>
        );
    }
}
