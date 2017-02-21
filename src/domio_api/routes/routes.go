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
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var RoutesList = Routes{
    Route{
        "CreateUser",
        "POST",
        "/users",
        create_user_handler.CreateUserHandler,
    },

    Route{
        "LoginUser",
        "POST",
        "/user/login",
        login_user_handler.LoginUserHandler,
    },

    Route{
        "GetDomainInfo",
        "GET",
        "/domain/{name}",
        get_domain_info_handler.GetDomainInfoHandler,
    },



    Route{
        "ShowStatus",
        "GET",
        "/",
        show_status_handler.ShowStatusHandler,
    },
    Route{
        "DeleteUser",
        "DELETE",
        "/user",
        delete_user_handler.DeleteUserHandler,
    },
    Route{
        "VerifyToken",
        "POST",
        "/tokens/verify",
        verify_token_handler.VerifyTokenHandler,
    },
    Route{
        "GetAvailableDomains",
        "GET",
        "/domains/available",
        get_available_domains_handler.GetAvailableDomainsHandler,
    },
    Route{
        "GetUserDomains",
        "GET",
        "/domains/user",
        get_user_domains_handler.GetUserDomainsHandler,
    },
    Route{
        "GetUserCards",
        "GET",
        "/cards",
        get_user_cards_handler.GetUserCardsHandler,
    },
    Route{
        "GetUserCard",
        "GET",
        "/cards/{id}",
        get_user_card_handler.GetUserCardHandler,
    },
    Route{
        "DeleteUserCard",
        "DELETE",
        "/cards/{id}",
        delete_card_handler.DeleteCardHandler,
    },
    Route{
        "CreateDomain",
        "POST",
        "/domains",
        create_domain_handler.CreateDomainHandler,
    },
    Route{
        "UpdateDomain",
        "PUT",
        "/domains/name/{name}",
        update_domain_handler.UpdateDomainHandler,
    },
    Route{
        "DeleteDomain",
        "DELETE",
        "/domain/{name}",
        delete_domain_handler.DeleteDomainHandler,
    },

    Route{
        "GetUserSubscriptions",
        "GET",
        "/subscriptions",
        get_user_subscriptions_handler.GetUserSubscriptionsHandler,
    },
    Route{
        "CreateSubscription",
        "POST",
        "/subscriptions",
        create_subscription_handler.CreateSubscriptionHandler,
    },
    Route{
        "UpdateSubscription",
        "PUT",
        "/subscriptions/{id}",
        update_subscription_handler.UpdateSubscriptionHandler,
    },
    Route{
        "GetSubscription",
        "GET",
        "/subscriptions/{id}",
        get_subscription_handler.GetSubscriptionHandler,
    },
    Route{
        "DeleteSubscription",
        "DELETE",
        "/subscriptions/{id}",
        delete_subscription_handler.DeleteSubscriptionHandler,
    },
    Route{
        "CreateCard",
        "POST",
        "/cards",
        create_card_handler.CreateCardHandler,
    },
}