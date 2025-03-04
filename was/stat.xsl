<?xml version="1.0"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
    <xsl:template match="/">
        <html>
        <head>
            <title>RTMP Streaming Status</title>
        </head>
        <body>
            <h1>RTMP Streaming Status</h1>
            <xsl:for-each select="rtmp/server/application">
                <h2>Application: <xsl:value-of select="name"/></h2>
                <xsl:if test="live/stream">
                    <xsl:for-each select="live/stream">
                        <h3>Stream Name: <xsl:value-of select="name"/></h3>
                        <p>Total clients: <xsl:value-of select="client_count"/></p>
                        <p>Time active: <xsl:value-of select="time"/></p>
                        <p>Bandwidth in: <xsl:value-of select="bw_in"/> Kbps</p>
                        <p>Bandwidth out: <xsl:value-of select="bw_out"/> Kbps</p>
                    </xsl:for-each>
                </xsl:if>
            </xsl:for-each>
        </body>
        </html>
    </xsl:template>
</xsl:stylesheet>