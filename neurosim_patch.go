package amqp

// NeurosimKeepsClientMaxMessageSize marks this patched build of go-amqp:
// v1.7.0 plus Azure/go-amqp#387, which keeps ReceiverOptions.MaxMessageSize
// when the peer omits max-message-size. NeurosimIO/neurosim-external-adapters
// references it so a build against unpatched go-amqp fails to compile.
const NeurosimKeepsClientMaxMessageSize = true
