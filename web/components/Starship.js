import React from 'react';
import {useQuery} from 'react-apollo';
import gql from "graphql-tag";
import {starship as starshipFragment} from "./fragments";
import {Spin, message, Col, Row, Button} from "antd";
import {useParams} from "react-router";
import {Link} from "react-router-dom";
import styles from './index.less';
import {Navs} from "./common";

export default function Starship() {
    const {id} = useParams();
    const {loading, error, data} = useQuery(gql`
        query {
            starship(id: ${id}) {
                ... StarshipParts
            }
        }
        ${starshipFragment}
    `);

    if (loading) {
        return <Spin spinning tip='loading starship info'/>
    }

    if (error) {
        message.error("load starship failure: " + error);
        return null;
    }

    const {starship} = data;
    const {name, length} = starship;

    return <Col span={12} offset={6}>
        <Row>
            Name: {name}
        </Row>
        <Row>
            Length: {length} feet
        </Row>
        <Navs/>
    </Col>
}

export function StarshipLink({className, id, name}) {
    return <Link className={className} to={`/starship/${id}`}>{name}</Link>
}
