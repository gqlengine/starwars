import React from 'react';
import { render } from 'react-dom';
import { HttpLink } from 'apollo-link-http'
import { ApolloLink } from 'apollo-link'
import ApolloClient from 'apollo-client'
import {InMemoryCache, IntrospectionFragmentMatcher} from 'apollo-cache-inmemory'
import { onError } from 'apollo-link-error'
import { createBrowserHistory } from 'history';
import {Route, Router, Switch} from "react-router";
import { ApolloProvider } from 'react-apollo';
import SearchList from "./components/SearchList";
import introspectionQueryResultData from './fragmentTypes';
import Human from "./components/Human";
import Droid from "./components/Droid";
import Starship from "./components/Starship";
import Episode from "./components/Episode";

const httpLink = new HttpLink({
    uri: 'http://localhost:9996/api/graphql',
});

const fragmentMatcher = new IntrospectionFragmentMatcher({
    introspectionQueryResultData,
});

export const cache = new InMemoryCache({
    fragmentMatcher,
});
export const history = createBrowserHistory();

const errLink = onError(
    ({ graphQLErrors, networkError, operation, forward }) => {
        if (graphQLErrors) {
            for (const err of graphQLErrors) {
                console.log(`[GraphQL error]: ${err}`)
            }
        }
        if (networkError) {
            console.log(`[Network error]: ${networkError}`)
            // if you would also like to retry automatically on
            // network errors, we recommend that you use
            // apollo-link-retry
        }
    }
);

export const apolloClient = new ApolloClient({
    cache,
    link: ApolloLink.from([
        errLink,
        httpLink,
    ]),
    connectToDevTools: true,
});


function App() {
    return (
        <Router history={history}>
            <ApolloProvider client={apolloClient}>
                <Switch>
                    <Route exact path='/episode/:episode'>
                        <Episode/>
                    </Route>
                    <Route exact path='/human/:id'>
                        <Human/>
                    </Route>
                    <Route exact path='/droid/:id'>
                        <Droid/>
                    </Route>
                    <Route exact path='/starship/:id'>
                        <Starship/>
                    </Route>
                    <Route>
                        <SearchList/>
                    </Route>
                </Switch>
            </ApolloProvider>
        </Router>
    )
}

render(<App />, document.querySelector('#root'));
