import React, { useCallback, useMemo } from "react";
import { useParams } from "react-router-dom";
import { FormattedMessage } from "@oursky/react-messageformat";

import { useAppConfigQuery } from "../portal/query/appConfigQuery";
import { useUserQuery } from "./query/userQuery";
import NavBreadcrumb from "../../NavBreadcrumb";
import { passwordFieldErrorRules } from "../../PasswordField";
import FormTextField from "../../FormTextField";
import AddIdentityForm from "./AddIdentityForm";
import ShowLoading from "../../ShowLoading";
import ShowError from "../../ShowError";
import { GenericErrorHandlingRule } from "../../error/useGenericError";

import styles from "./AddUsernameScreen.module.scss";

interface UsernameFieldProps {
  value: string;
  onChange: (value: string) => void;
}

const UsernameField: React.FC<UsernameFieldProps> = function UsernameField(
  props
) {
  const { value, onChange } = props;
  const onUsernameChange = useCallback(
    (_, value?: string) => onChange(value ?? ""),
    [onChange]
  );
  return (
    <FormTextField
      parentJSONPointer=""
      jsonPointer="username"
      fieldName="username"
      fieldNameMessageID="AddUsernameScreen.username.label"
      className={styles.usernameField}
      value={value}
      onChange={onUsernameChange}
    />
  );
};

const AddUsernameScreen: React.FC = function AddUsernameScreen() {
  const { appID, userID } = useParams();
  const {
    user,
    loading: loadingUser,
    error: userError,
    refetch: refetchUser,
  } = useUserQuery(userID);
  const {
    effectiveAppConfig,
    loading: loadingAppConfig,
    error: appConfigError,
    refetch: refetchAppConfig,
  } = useAppConfigQuery(appID);

  const navBreadcrumbItems = useMemo(() => {
    return [
      { to: "../../..", label: <FormattedMessage id="UsersScreen.title" /> },
      { to: "..", label: <FormattedMessage id="UserDetailsScreen.title" /> },
      { to: ".", label: <FormattedMessage id="AddUsernameScreen.title" /> },
    ];
  }, []);
  const title = <NavBreadcrumb items={navBreadcrumbItems} />;

  const rules: GenericErrorHandlingRule[] = useMemo(
    () => [
      {
        reason: "InvariantViolated",
        kind: "DuplicatedIdentity",
        errorMessageID: "AddUsernameScreen.error.duplicated-username",
        field: "username",
      },
      ...passwordFieldErrorRules,
    ],
    []
  );

  if (loadingUser || loadingAppConfig) {
    return <ShowLoading />;
  }

  if (userError != null) {
    return <ShowError error={userError} onRetry={refetchUser} />;
  }

  if (appConfigError != null) {
    return <ShowError error={appConfigError} onRetry={refetchAppConfig} />;
  }

  return (
    <AddIdentityForm
      appConfig={effectiveAppConfig}
      rawUser={user}
      loginIDType="username"
      title={title}
      loginIDField={UsernameField}
      errorRules={rules}
    />
  );
};

export default AddUsernameScreen;
