import gql from "graphql-tag";

export const human = gql`
    fragment HumanParts on Human {
        id
        name
        height(unit: FOOT)
        mass
        appearsIn
        starships {
            id
            name
        }
        friends {
            id
            name
        }
    }
`;

export const droid = gql`
    fragment DroidParts on Droid {
        id
        name
        appearsIn
        primaryFunction
        friends {
            id
            name
        }
    }
`;

export const starship = gql`
    fragment StarshipParts on Starship {
        id
        name
        length
    }
`;
