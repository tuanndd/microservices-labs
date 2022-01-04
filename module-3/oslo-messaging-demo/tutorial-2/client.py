import sys
import threading
import time

from oslo_log import log as logging
from oslo_config import cfg
import oslo_messaging as messaging

LOG = logging.getLogger(__name__)

_options = [
    cfg.StrOpt("topic",
               default="my-topic",
               help="target topic, default 'my-topic'"),
    cfg.StrOpt('exchange',
               default="my-exchange",
               help="target exchange, default 'my-exchange'"),
    cfg.StrOpt('namespace',
               help="target namespace, default None"),
    cfg.StrOpt("url",
               required=True,
               default="rabbit://localhost:5672",
               help="transport address, default 'rabbit://localhost:5672'"),
    cfg.StrOpt("target_version",
               help="Override the default version in the target address"),
    cfg.BoolOpt("quiet",
                default=False,
                help="Suppress all stdout output"),
    cfg.StrOpt("server",
               help="Send only to the named server"),
    cfg.IntOpt("timeout",
               default=60,
               help="timeout RPC request in seconds, default 60"),
    cfg.BoolOpt("fanout",
                default=False,
                help="Send RPC fanout cast"),           
    cfg.StrOpt("log_levels",
               help="Set module specific log levels, e.g."
               " 'amqp=WARN,oslo.messaging=INFO,...'")
]

def main(argv=None):

    logging.register_options(cfg.CONF)
    cfg.CONF.register_cli_opts(_options)
    cfg.CONF(sys.argv[1:])
    if cfg.CONF.log_levels:
        logging.set_defaults(
            default_log_levels=cfg.CONF.log_levels.split(',')
        )
    logging.setup(cfg.CONF, "client")

    quiet = cfg.CONF.quiet
    server = cfg.CONF.server
    exchange = cfg.CONF.exchange
    topic = cfg.CONF.topic
    namespace = cfg.CONF.namespace
    url = cfg.CONF.url
    target_version = cfg.CONF.target_version
    fanout = cfg.CONF.fanout

    transport = messaging.get_rpc_transport(cfg.CONF, url=url)

    target = messaging.Target(exchange=exchange,
                              topic=topic,
                              namespace=namespace,
                              server=server,
                              fanout=fanout,
                              version=target_version)
    
    client = messaging.RPCClient(transport,
                                     target,
                                     timeout=cfg.CONF.timeout,
                                     version_cap=cfg.CONF.target_version).prepare()
    
    # RPC Call (***not support fanout***)
    # arg = "Saju"
    # ctx = {}
    # resp = client.call(ctx, 'test_method1', arg=arg)
    # print("RPC return value=%s" % str(resp))

    # RPC Cast, fanout
    ctx = {}
    for x in range(10):
        client.cast(ctx, 'test_method1', arg=x)

if __name__ == "__main__":
    sys.exit(main())
