import React, {forwardRef} from 'react';
import {useQuery} from 'react-apollo/lib/index';
import gql from "graphql-tag";
import {Select, Tooltip} from "antd";

const {Option} = Select;

const EnumSelector = forwardRef(({enumName, ...props}, ref) => {
    const {loading, error, data} = useQuery(gql`{
        enumDef: __type(name: "${enumName}") {
            name
            kind
            description
            enumValues {
                name
                description
                isDeprecated
                deprecationReason
            }
        }
    }`);

    if (loading) {
        return <Select {...props} loading={loading}/>
    }

    if (error) {
        return <Select {...props} disabled={true}/>
    }

    const {enumDef} = data;
    const {name, description, enumValues} = enumDef;
    return (
        <div ref={ref}>
            <Select {...props} placeholder={props.placeholder || description}>
                {enumValues.map(({name, description, isDeprecated, deprecatedReason}) => (
                    <Option key={name} value={name} disabled={isDeprecated}>
                        {isDeprecated&&(
                            <Tooltip title={deprecatedReason}>
                                <span>{description}</span>
                            </Tooltip>
                        )}
                        {!isDeprecated&&description}
                    </Option>
                ))}
            </Select>
        </div>
    )
});

export default EnumSelector;
