import React, {useState} from 'react';
import {useQuery} from 'react-apollo';
import gql from "graphql-tag";
import {Input, Layout, Row, Typography, List, Affix, Col} from 'antd';
import {droid, human, starship} from "./fragments";
import {Link} from "react-router-dom";
import styles from './index.less';
import {HumanLink} from "./Human";
import {DroidLink} from "./Droid";
import {StarshipLink} from "./Starship";
import {EpisodeLink} from "./Episode";
import FriendLink, {Navs} from "./common";

const {Search} = Input;
const {Text} = Typography;
const {Item} = List;

export default function SearchList({}) {
    const [searchText, setSearchText] = useState('');
    const {loading, error, data} = useQuery(gql`
        query {
            results: search(text: "${searchText}") {
                __typename
                ... HumanParts
                ... DroidParts
                ... StarshipParts 
            }
        }
        
        ${human}
        ${droid}
        ${starship}
    `);

    const search = (e) => {
        e.preventDefault();
        setSearchText(e.target.value);
    };

    const {results} = data || {};

    console.log(results);

    const renderItem = (result) => {
        switch (result.__typename) {
            case 'Human':
                return <HumanItem {...result}/>;
            case 'Droid':
                return <DroidItem {...result}/>;
            case 'Starship':
                return <StarshipItem {...result}/>
        }
    };

    return <Col span={12} offset={6} style={{paddingTop: 10}}>
        <Affix offsetTop={10}>
            <Search placeholder='input text to search peoples/starships about starwar'
                    value={searchText}
                    onChange={search}
                    loading={loading}
            />
        </Affix>
        <Row>
            <List  itemLayout="horizontal"
                   dataSource={results && results.filter(r => !!r)}
                   renderItem={renderItem}/>
        </Row>
        <Navs playground/>
    </Col>
}

function HumanItem({id, name, height, mass, appearsIn, friends, starships}) {
    return <Item key={id}>
        <Item.Meta
            title={<Text strong><HumanLink id={id} name={name}/></Text>}
            description={<div>
                <div>height: {height} feet</div>
                <div>mass: {mass} kilograms</div>
                <div>
                    Appears in:
                    {appearsIn&&appearsIn.filter(r => !!r).map(episode => (
                        <Text key={episode} code className={styles.gap}>
                            <EpisodeLink episode={episode}/>
                        </Text>
                    ))}
                </div>
                <div>
                    Friends:
                    {friends&&friends.filter(r => !!r).map((friend) => (
                        <Text key={friend.id} code className={styles.gap}>
                            <FriendLink {...friend}/>
                        </Text>
                    ))}
                </div>
                <div>
                    Starships:
                    {starships&&starships.filter(r => !!r).map(({id, name}) => (
                        <Text key={id} code className={styles.gap}>
                            <StarshipLink id={id} name={name}/>
                        </Text>
                    ))}
                </div>
            </div>}
        />
    </Item>
}

function DroidItem({id, name, primaryFunction, appearsIn, friends}) {
    return <Item key={id}>
        <Item.Meta
            title={<Text strong><DroidLink id={id} name={name}/></Text>}
            description={<div>
                <div>Primary function: {primaryFunction}</div>
                <div>
                    Appears in:
                    {appearsIn&&appearsIn.filter(r => !!r).map(episode => (
                        <Text key={episode} code className={styles.gap}>
                            <EpisodeLink episode={episode}/>
                        </Text>
                    ))}
                </div>
                <div>
                    Friends:
                    {friends&&friends.filter(r => !!r).map((friend) => (
                        <Text key={friend.id} code className={styles.gap}>
                            <FriendLink {...friend}/>
                        </Text>
                    ))}
                </div>
            </div>}
        />
    </Item>
}

function StarshipItem({id, name, length}) {
    return <Item key={id}>
        <Item.Meta
            title={<Text strong><StarshipLink id={id} name={name}/></Text>}
            description={<Text>length: {length} feet</Text>}
        />
    </Item>
}
