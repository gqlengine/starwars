import React from 'react';
import {HumanLink} from "./Human";
import {DroidLink} from "./Droid";
import {Link} from "react-router-dom";
import styles from './index.less';

export default function FriendLink({__typename, className, id, name}) {
    switch (__typename) {
        case 'Human':
            return <HumanLink className={className} id={id} name={name}/>;
        case 'Droid':
            return <DroidLink className={className} id={id} name={name}/>;
    }
    return null;
}

export function Navs({playground}) {
    return <div className={styles.goBackToList}>
        {!playground && (
            <>
                <Link to='/'>return to list</Link>
                <br/>
            </>
        )}
        <Link to='/api/graphql/playground/' target='_blank'>open Playground</Link>
    </div>
}

