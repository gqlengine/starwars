import React from 'react';
import {Link, useParams} from "react-router-dom";
import gql from "graphql-tag";
import {useQuery, useMutation} from 'react-apollo';
import {droid, human} from "./fragments";
import {Button, Col, Divider, Form, Input, List, message, Rate, Row, Typography, Spin} from "antd";
import FriendLink, {Navs, OpenPlayground} from "./common";
import EnumSelector from './EnumSelector'

const {Text} = Typography;
import styles from './index.less';

export function EpisodeLink({className, episode}) {
    return <Link className={className} to={`/episode/${episode}`}>{episode}</Link>
}

export default function Episode() {
    const {episode} = useParams();
    const {loading, error, data} = useQuery(gql`
        query EpisodeQuery {
            hero(episode: ${episode}) {
                ... HumanParts
                ... DroidParts
            }
            reviews(episode: ${episode}) {
                stars
                commentary
            }
        }
        ${human}
        ${droid}
    `);

    if (loading) {
        return <Spin spinning tip='loading episode info'/>
    }

    if (error) {
        message.error("load episode failure: " + error);
        return null;
    }

    const {reviews, hero} = data;

    return <Col span={12} offset={6}>
        <Row>
            Current is <Text strong>{episode}</Text>, other episodes:
            {episode !== 'NEWHOPE' && <EpisodeLink className={styles.gap} episode='NEWHOPE'/>}
            {episode !== 'EMPIRE' && <EpisodeLink className={styles.gap} episode='EMPIRE'/>}
            {episode !== 'JEDI' && <EpisodeLink className={styles.gap} episode='JEDI'/>}
        </Row>
        <Row>
            Hero: {hero && <FriendLink {...hero}/>}
        </Row>
        <Row>
            Reviews:
            <List>
                {reviews&&reviews.map(({stars, commentary}, index) => (
                    <List.Item key={index}>
                        <List.Item.Meta
                            title={<span>
                                #{index} - <Rate disabled defaultValue={stars}/>
                            </span>}
                            description={commentary}
                        />
                    </List.Item>
                ))}
            </List>
        </Row>
        <Divider> Comment this episode! </Divider>
        <Row>
            <Col span={12}>
                <Row>
                    <CommentPanel/>
                </Row>
            </Col>
        </Row>
        <Navs/>
    </Col>
}

function Comment(
    {form: {getFieldDecorator, validateFields, resetFields}}
) {
    const {episode} = useParams();
    const [createReview, {loading, error}] = useMutation(gql`
    mutation (
        $episode: Episode!
        $stars: Int!
        $commentary: String
    ){
        createReviews(
            episode: $episode
            review: {
                stars: $stars
                commentary: $commentary
            }
        ) {stars}
    }`, {
        refetchQueries: ['EpisodeQuery']
    });

    const submit = (e) => {
        e.preventDefault();

        validateFields((err, variables) => {
            if (err) {
                message.error(err);
                return;
            }

            createReview({
                variables
            }).then(() => {
                resetFields()
            }).catch((err) => {
                message.error(err);
            })
        })
    };

    return <Form onSubmit={submit}>
        <Form.Item label='Choose episode'>
            {getFieldDecorator('episode', {
                initialValue: episode,
                rules: [{
                    required: true,
                    message: 'please choose episode',
                }]
            })(
                <EnumSelector enumName='Episode'/>
            )}
        </Form.Item>
        <Form.Item label='Rate'>
            {getFieldDecorator('stars', {
                rules: [{
                    required: true,
                    message: 'please (re)rate episode',
                }]
            })(
                <Rate/>
            )}
        </Form.Item>
        <Form.Item label='Comments'>
            {getFieldDecorator('commentary')(
                <Input.TextArea placeholder={'comment the episode here'}/>
            )}
        </Form.Item>
        <Form.Item>
            <Button width='100%' type='primary' htmlType='submit'
                    loading={loading}>
                Comment
            </Button>
        </Form.Item>
    </Form>
}

export const CommentPanel = Form.create({name: 'review'})(Comment);
