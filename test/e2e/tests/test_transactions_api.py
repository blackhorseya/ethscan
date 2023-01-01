from typing import Generator

import pytest
from playwright.sync_api import Playwright, APIRequestContext

BASE_URL = "https://ethscan.seancheng.space"
TX_HASH = ""


@pytest.fixture(scope="session")
def api_request_context(playwright: Playwright) -> Generator[APIRequestContext, None, None]:
    headers = {
        "Accept": "application/json",
    }
    request_context = playwright.request.new_context(base_url=BASE_URL, extra_http_headers=headers)

    yield request_context

    request_context.dispose()


@pytest.fixture(scope="session", autouse=True)
def get_tx_hash(api_request_context: APIRequestContext) -> None:
    resp = api_request_context.get("/api/v1/blocks")
    assert resp.ok

    body = resp.json()
    block = body["data"]["list"][0]
    assert block

    block_hash = block["block_hash"]

    resp = api_request_context.get(f"/api/v1/blocks/{block_hash}")
    assert resp.ok

    body = resp.json()
    block = body["data"]
    assert block

    global TX_HASH
    TX_HASH = block["transactions"][0]


def test_should_get_tx_by_hash(api_request_context: APIRequestContext) -> None:
    resp = api_request_context.get(f"/api/v1/transactions/{TX_HASH}")
    assert resp.ok

    body = resp.json()
    tx = body["data"]
    assert tx
