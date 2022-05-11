import React from 'react';
import api from '../core/api-service.js';
import axios from 'axios';
import VeloTable from '../core/table.js';
import FormControl from 'react-bootstrap/FormControl'
import Button from 'react-bootstrap/Button'
import InputGroup from 'react-bootstrap/InputGroup'
import FormGroup from 'react-bootstrap/FormGroup'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

export default class UserPage extends React.Component {

    componentDidMount = () => {
        this.source = axios.CancelToken.source();
        this.loadUsers();
    }

    componentWillUnmount() {
        this.source.cancel("unmounted");
    }

    componentDidUpdate = (prevProps, prevState, rootNode) => {
        if (!this.state.initialized) {
            this.loadUsers();
        }
    }

    state = {
        users: [],
        columns: ["Username", "DeleteUser"],
        initialized: false,
        createUsername: "",
        createPassword: "",
    }

    handleDelete = (username, e) => {
        e.preventDefault();
        var result = window.confirm("Do you want to delete " + username + "?")
        if (result) {
            api.delete_req("v1/DeleteUser/" + username, {}, this.source.token)
            this.loadUsers();
        }
    }

    handleCreate = (e) => {
        e.preventDefault();

        if (this.state.createUsername !== "" && this.state.createPassword !== "") {
            api.post("v1/CreateUser", {name: this.state.createUsername, password: this.state.createPassword}, this.source.token)
            this.loadUsers();
        }
    }

    loadUsers = () => {
        api.get("v1/GetUsers", {}, this.source.token).then((response) => {
            let names = [];
            for(var i = 0; i<response.data.users.length; i++) {
                var name = response.data.users[i].name;

                names.push({Username: name, DeleteUser: name});
            }
            this.setState({users: names, initialized: true});
        });
    };

    render() {
        let renderers = {
            DeleteUser: (cell, row) => {
                return <Button id="delete_user"
                               onClick={(e) => this.handleDelete(cell, e)}
                               variant="danger">
                         <FontAwesomeIcon icon="window-close"/>
                       </Button>
            },
        };
        return(
            <>
            <div className="col-12">
            <h2>Create a new user</h2>
            <FormGroup>
            <InputGroup>
                <FormControl
                  placeholder="Username"
                  aria-label="Username"
                  aria-describedby="basic-addon1"
                  onChange={(e) => this.setState({createUsername: e.target.value})}
                />
                <FormControl
                  placeholder="Password"
                  aria-label="Password"
                  type="password"
                  aria-describedby="basic-addon1"
                  onChange={(e) => this.setState({createPassword: e.target.value})}
                />
                <Button id="create_user_submit"
            onClick={this.handleCreate}
            variant="default" type="submit">
                Add User
            </Button>
            </InputGroup>
            </FormGroup>
            </div>
            <div className="col-12">
            <h2>Users</h2>
            </div>
            <VeloTable
            rows={this.state.users}
            columns={this.state.columns}
            renderers={renderers}
            />
                </>
        )
    }
};
