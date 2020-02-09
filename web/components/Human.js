import React from 'react';
import {useQuery} from 'react-apollo';
import gql from "graphql-tag";
import {human as humanFragment} from "./fragments";
import {Spin, message, Col, Row, Button} from "antd";
import {useParams} from "react-router";
import {Link} from "react-router-dom";
import styles from "./index.less";
import {StarshipLink} from "./Starship";
import {DroidLink} from "./Droid";
import {Navs} from "./common";

export default function Human() {
    const {id} = useParams();
    const {loading, error, data} = useQuery(gql`
        query {
            human(id: ${id}) {
                ... HumanParts
            }
        }
        ${humanFragment}
    `);

    if (loading) {
        return <Spin spinning tip='loading human info'/>
    }

    if (error) {
        message.error("load human failure: " + error);
        return null;
    }

    const {human} = data;
    const {name, height, mass, friends, appearsIn, starships} = human;

    return <Col span={12} offset={6}>
        <Row>
            Name: {name}
        </Row>
        <Row>
            Height: {height} feet
        </Row>
        <Row>
            Mass: {mass} kilograms
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
        <Row>
            Starships:
            {starships&&starships.filter(r => !!r).map((starship) => {
                return <StarshipLink className={styles.gap} {...starship}/>
            })}
        </Row>
        <Navs/>
    </Col>
}

export function HumanLink({className, id, name}) {
    return <Link className={className} to={`/human/${id}`}>{name}</Link>
}
