package routes

import (
    "net/http"
    "domio_api/handlers/create_user_handler"
    "domio_api/handlers/login_user_handler"
    "domio_api/handlers/verify_token_handler"
    "domio_api/handlers/get_available_domains_handler"
    "domio_api/handlers/get_user_domains_handler"
    "domio_api/handlers/create_domain_handler"
    "domio_api/handlers/get_domain_info_handler"
    "domio_api/handlers/create_subscription_handler"
    "domio_api/handlers/create_card_handler"
    "domio_api/handlers/delete_domain_handler"
    "domio_api/handlers/delete_subscription_handler"
    "domio_api/handlers/get_user_subscriptions_handler"
    "domio_api/handlers/show_status_handler"
    "domio_api/handlers/delete_user_handler"
    "domio_api/handlers/update_domain_handler"
    "domio_api/handlers/get_user_cards_handler"
    "domio_api/handlers/get_user_card_handler"
    "domio_api/handlers/delete_card_handler"
    "domio_api/handlers/get_subscription_handler"
    "domio_api/handlers/update_subscription_handler"
    "domio_api/handlers/get_subscription_records_handler"
    "domio_api/handlers/update_subscription_records_handler"
    "domio_api/handlers/delete_record_handler"
)

type CheckAccessFunc func(w http.ResponseWriter) bool

const (
    Nobody = 0
    Everyone = 1
    User = 2
    Owner = 3
    Administrator = 4
)

type Route struct {
    Name            string
    Method          string
    Pattern         string
    HandlerFunc     http.HandlerFunc
    CheckAccessFunc CheckAccessFunc
}

type Routes []Route

var RoutesList = Routes{
    Route{
        "CreateUser",
        http.MethodPost,
        "/users",
        create_user_handler.CreateUserHandler,
        Everyone,
    },

    Route{
        "LoginUser",
        http.MethodPost,
        "/user/login",
        login_user_handler.LoginUserHandler,
        login_user_handler.CheckAccessFunc,
    },

    Route{
        "GetDomainInfo",
        http.MethodGet,
        "/domain/{name}",
        get_domain_info_handler.GetDomainInfoHandler,
        Everyone,
    },

    Route{
        "ShowStatus",
        http.MethodGet,
        "/",
        show_status_handler.ShowStatusHandler,
        Everyone,
    },

    Route{
        "DeleteUser",
        http.MethodDelete,
        "/user",
        delete_user_handler.DeleteUserHandler,
        Owner,
    },

    Route{
        "VerifyToken",
        http.MethodPost,
        "/tokens/verify",
        verify_token_handler.VerifyTokenHandler,
        Everyone,
    },

    Route{
        "GetAvailableDomains",
        http.MethodGet,
        "/domains/available",
        get_available_domains_handler.GetAvailableDomainsHandler,
        Everyone,
    },

    Route{
        "GetUserDomains",
        http.MethodGet,
        "/user/domains",
        get_user_domains_handler.GetUserDomainsHandler,
        Owner,
    },

    Route{
        "GetUserCards",
        http.MethodGet,
        "/cards",
        get_user_cards_handler.GetUserCardsHandler,
        Owner,
    },

    Route{
        "GetUserCard",
        http.MethodGet,
        "/cards/{id}",
        get_user_card_handler.GetUserCardHandler,
        Owner,
    },

    Route{
        "DeleteUserCard",
        http.MethodDelete,
        "/cards/{id}",
        delete_card_handler.DeleteCardHandler,
        Owner,
    },
    Route{
        "CreateDomain",
        http.MethodPost,
        "/domains",
        create_domain_handler.CreateDomainHandler,
        User,
    },

    Route{
        "UpdateDomain",
        http.MethodPut,
        "/domain/{name}",
        update_domain_handler.UpdateDomainHandler,
        Owner,
    },

    Route{
        "DeleteDomain",
        http.MethodDelete,
        "/domain/{name}",
        delete_domain_handler.DeleteDomainHandler,
        Owner,
    },

    Route{
        "GetUserSubscriptions",
        http.MethodGet,
        "/subscriptions",
        get_user_subscriptions_handler.GetUserSubscriptionsHandler,
        Owner,
    },

    Route{
        "CreateSubscription",
        http.MethodPost,
        "/subscriptions",
        create_subscription_handler.CreateSubscriptionHandler,
        User,
    },

    Route{
        "UpdateSubscription",
        http.MethodPut,
        "/subscriptions/{id}",
        update_subscription_handler.UpdateSubscriptionHandler,
        Owner,
    },

    Route{
        "GetSubscription",
        http.MethodGet,
        "/subscriptions/{id}",
        get_subscription_handler.GetSubscriptionHandler,
        Owner,
    },

    Route{
        "DeleteSubscription",
        http.MethodDelete,
        "/subscription/{subId}",
        delete_subscription_handler.DeleteSubscriptionHandler,
        Owner,
    },

    Route{
        "DeleteRecord",
        http.MethodDelete,
        "/subscription/{subId}/records",
        delete_record_handler.DeleteRecordHandler,
        Owner,
    },

    Route{
        "CreateCard",
        http.MethodPost,
        "/cards",
        create_card_handler.CreateCardHandler,
        User,
    },

    Route{
        "GetSubscriptionRecords",
        http.MethodGet,
        "/subscriptions/{id}/records",
        get_subscription_records_handler.GetSubscriptionRecordsHandler,
        Owner,
    },

    Route{
        "UpdateSubscriptionRecords",
        http.MethodPut,
        "/subscription/{id}/records",
        update_subscription_records_handler.UpdateSubscriptionRecordsHandler,
        Owner,
    },
}