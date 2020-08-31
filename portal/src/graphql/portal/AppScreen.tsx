import React from "react";
import { graphql, QueryRenderer } from "react-relay";
import { useParams } from "react-router-dom";
import { AppScreenQueryResponse } from "./__generated__/AppScreenQuery.graphql";
import { environment } from "./relay";
import UserList from "../adminapi/UserList";
import ShowError from "../../ShowError";
import ShowLoading from "../../ShowLoading";

const query = graphql`
  query AppScreenQuery($id: ID!) {
    node(id: $id) {
      ... on App {
        id
        appConfig
        secretConfig
      }
    }
  }
`;

interface Variables {
  id: string;
}

const ShowApp: React.FC<AppScreenQueryResponse> = function ShowApp(
  props: AppScreenQueryResponse
) {
  return (
    <div>
      <pre>{JSON.stringify(props.node, null, 2)}</pre>
      <br />
      <UserList />
    </div>
  );
};

const AppScreen: React.FC = function AppScreen() {
  const { appID } = useParams();
  return (
    <QueryRenderer<{ variables: Variables; response: AppScreenQueryResponse }>
      environment={environment}
      query={query}
      variables={{ id: appID }}
      render={({ error, props, retry }) => {
        if (error != null) {
          return <ShowError error={error} onRetry={retry} />;
        }

        if (props == null) {
          return <ShowLoading />;
        }
        return <ShowApp {...props} />;
      }}
    />
  );
};

export default AppScreen;