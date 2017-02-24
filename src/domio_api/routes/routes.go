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

const (
    Everyone = 0
    User = 1
    Owner = 2
    Administrator = 3
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
    AccessLevel uint64
}

type Routes []Route

var RoutesList = Routes{
    Route{
        "CreateUser",
        "POST",
        "/users",
        create_user_handler.CreateUserHandler,
        Everyone,
    },

    Route{
        "LoginUser",
        "POST",
        "/user/login",
        login_user_handler.LoginUserHandler,
        Everyone,
    },

    Route{
        "GetDomainInfo",
        "GET",
        "/domain/{name}",
        get_domain_info_handler.GetDomainInfoHandler,
        Everyone,
    },

    Route{
        "ShowStatus",
        "GET",
        "/",
        show_status_handler.ShowStatusHandler,
        Everyone,
    },

    Route{
        "DeleteUser",
        "DELETE",
        "/user",
        delete_user_handler.DeleteUserHandler,
        Owner,
    },

    Route{
        "VerifyToken",
        "POST",
        "/tokens/verify",
        verify_token_handler.VerifyTokenHandler,
        Everyone,
    },

    Route{
        "GetAvailableDomains",
        "GET",
        "/domains/available",
        get_available_domains_handler.GetAvailableDomainsHandler,
        Everyone,
    },

    Route{
        "GetUserDomains",
        "GET",
        "/user/domains",
        get_user_domains_handler.GetUserDomainsHandler,
        Owner,
    },

    Route{
        "GetUserCards",
        "GET",
        "/cards",
        get_user_cards_handler.GetUserCardsHandler,
        Owner,
    },

    Route{
        "GetUserCard",
        "GET",
        "/cards/{id}",
        get_user_card_handler.GetUserCardHandler,
        Owner,
    },

    Route{
        "DeleteUserCard",
        "DELETE",
        "/cards/{id}",
        delete_card_handler.DeleteCardHandler,
        Owner,
    },
    Route{
        "CreateDomain",
        "POST",
        "/domains",
        create_domain_handler.CreateDomainHandler,
        User,
    },

    Route{
        "UpdateDomain",
        "PUT",
        "/domain/{name}",
        update_domain_handler.UpdateDomainHandler,
        Owner,
    },

    Route{
        "DeleteDomain",
        "DELETE",
        "/domain/{name}",
        delete_domain_handler.DeleteDomainHandler,
        Owner,
    },

    Route{
        "GetUserSubscriptions",
        "GET",
        "/subscriptions",
        get_user_subscriptions_handler.GetUserSubscriptionsHandler,
        Owner,
    },

    Route{
        "CreateSubscription",
        "POST",
        "/subscriptions",
        create_subscription_handler.CreateSubscriptionHandler,
        User,
    },

    Route{
        "UpdateSubscription",
        "PUT",
        "/subscriptions/{id}",
        update_subscription_handler.UpdateSubscriptionHandler,
        Owner,
    },

    Route{
        "GetSubscription",
        "GET",
        "/subscriptions/{id}",
        get_subscription_handler.GetSubscriptionHandler,
        Owner,
    },

    Route{
        "DeleteSubscription",
        "DELETE",
        "/subscription/{subId}",
        delete_subscription_handler.DeleteSubscriptionHandler,
        Owner,
    },

    Route{
        "DeleteRecord",
        "DELETE",
        "/subscription/{subId}/records",
        delete_record_handler.DeleteRecordHandler,
        Owner,
    },

    Route{
        "CreateCard",
        "POST",
        "/cards",
        create_card_handler.CreateCardHandler,
        User,
    },

    Route{
        "GetSubscriptionRecords",
        "GET",
        "/subscriptions/{id}/records",
        get_subscription_records_handler.GetSubscriptionRecordsHandler,
        Owner,
    },

    Route{
        "UpdateSubscriptionRecords",
        "PUT",
        "/subscription/{id}/records",
        update_subscription_records_handler.UpdateSubscriptionRecordsHandler,
        Owner,
    },
}