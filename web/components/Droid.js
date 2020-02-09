import React from 'react';
import {useQuery} from 'react-apollo';
import gql from "graphql-tag";
import {droid as droidFragment} from "./fragments";
import {Spin, message, Col, Row, Button, Icon} from "antd";
import {useParams} from "react-router";
import {Link} from "react-router-dom";
import styles from './index.less';
import {HumanLink} from "./Human";
import {Navs} from "./common";

export default function Droid() {
    const {id} = useParams();
    const {loading, error, data} = useQuery(gql`
        query {
            droid(id: ${id}) {
                ... DroidParts
            }
        }
        ${droidFragment}
    `);

    if (loading) {
        return <Spin spinning tip='loading droid info'/>
    }

    if (error) {
        message.error("load droid failure: " + error);
        return null;
    }

    const {droid} = data;
    const {name, friends, appearsIn} = droid;

    return <Col span={12} offset={6}>
        <Row>
            Name: {name}
        </Row>
        <Row>
            Friends:
            {friends&&friends.filter(r => !!r).map((friend) => {
                switch (friend.__typename) {
                    case 'Human':
                        return <HumanLink className={styles.gap} key={friend.id}/>
                    case 'Droid':
                        return <DroidLink className={styles.gap} key={friend.id}/>
                }
            })}
        </Row>
        <Navs/>
    </Col>
}

export function DroidLink({className, id, name}) {
    return <Link className={className} to={`/droid/${id}`}>{name}</Link>
}
