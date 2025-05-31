#!/bin/sh

if [ "X$1" = "X" ]; then
    printf "Usage: mkdir.sh dirname";
    exit
fi

mkdir -p "$1/domain"        # ドメイン(エンティティ, バリューオブジェクト)
mkdir -p "$1/usecase"       # アプリケーションサービス(サービス=ユースケース)
mkdir -p "$1/repository"    # インフラ(リポジトリ, 外部サービス)
mkdir -p "$1/presentation"  # プレゼンテーション層(APIハンドラ, ルータ)
