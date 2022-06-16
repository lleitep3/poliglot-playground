FROM php:8.0.20-zts-alpine3.16

RUN curl -sS https://getcomposer.org/installer | php \
    && mv composer.phar /usr/local/bin/composer

RUN apk add --no-cache libzip-dev $PHPIZE_DEPS \
    && pecl install xdebug \
    && docker-php-ext-enable xdebug \
    && docker-php-ext-configure zip \
    && docker-php-ext-install zip

WORKDIR /code

COPY . /code

CMD ["composer", "tests"]
