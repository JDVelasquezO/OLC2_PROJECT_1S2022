// Generated from /home/jdvelasquezo/GolandProjects/OLC2_Project1/server/ProjectLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ProjectLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PRINTLN=1, PRINT=2, DECLARAR=3, RSTRING=4, RINTEGER=5, RREAL=6, RBOOLEAN=7, 
		RSTR=8, RIF=9, RELSE=10, RAND=11, ROR=12, MULTILINE=13, INLINE=14, INTEGER=15, 
		FLOAT=16, CHAR=17, STRING=18, BOOLEAN=19, ID=20, EQUAL=21, DOT=22, SEMICOLON=23, 
		COMMA=24, COLON=25, ADMIRATION=26, REFERENCE=27, LEFT_PAR=28, RIGHT_PAR=29, 
		LEFT_KEY=30, RIGHT_KEY=31, LEFT_BRACKET=32, RIGHT_BRACKET=33, EQUEAL_EQUAL=34, 
		NOTEQUAL=35, GREATER_THAN=36, LESS_THAN=37, GREATER_EQUALTHAN=38, LESS_EQUEALTHAN=39, 
		MUL=40, DIV=41, ADD=42, SUB=43, MOD=44, WHITESPACE=45;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PRINTLN", "PRINT", "DECLARAR", "RSTRING", "RINTEGER", "RREAL", "RBOOLEAN", 
			"RSTR", "RIF", "RELSE", "RAND", "ROR", "MULTILINE", "INLINE", "INTEGER", 
			"FLOAT", "CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", "SEMICOLON", 
			"COMMA", "COLON", "ADMIRATION", "REFERENCE", "LEFT_PAR", "RIGHT_PAR", 
			"LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", 
			"NOTEQUAL", "GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", 
			"MUL", "DIV", "ADD", "SUB", "MOD", "WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println'", "'print'", "'let'", "'String'", "'Integer'", "'Real'", 
			"'Boolean'", "'str'", "'If'", "'entonces'", "'AND'", "'OR'", null, null, 
			null, null, null, null, null, null, "'='", "'.'", "';'", "','", "':'", 
			"'!'", "'&'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'=='", "'!='", 
			"'>'", "'<'", "'>='", "'<='", "'*'", "'/'", "'+'", "'-'", "'%'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PRINTLN", "PRINT", "DECLARAR", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RSTR", "RIF", "RELSE", "RAND", "ROR", "MULTILINE", "INLINE", 
			"INTEGER", "FLOAT", "CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", 
			"SEMICOLON", "COMMA", "COLON", "ADMIRATION", "REFERENCE", "LEFT_PAR", 
			"RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET", 
			"EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN", 
			"LESS_EQUEALTHAN", "MUL", "DIV", "ADD", "SUB", "MOD", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public ProjectLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ProjectLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2/\u0128\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t"+
		"\3\t\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3\16\3\16\6\16\u00a9\n\16\r\16"+
		"\16\16\u00aa\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\6\17\u00b6\n"+
		"\17\r\17\16\17\u00b7\3\17\3\17\3\20\6\20\u00bd\n\20\r\20\16\20\u00be\3"+
		"\21\6\21\u00c2\n\21\r\21\16\21\u00c3\3\21\3\21\6\21\u00c8\n\21\r\21\16"+
		"\21\u00c9\3\22\3\22\3\22\3\22\3\23\3\23\7\23\u00d2\n\23\f\23\16\23\u00d5"+
		"\13\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u00e2"+
		"\n\24\3\25\3\25\7\25\u00e6\n\25\f\25\16\25\u00e9\13\25\3\26\3\26\3\27"+
		"\3\27\3\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\36"+
		"\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3#\3$\3$\3$\3%\3%\3&\3&\3\'"+
		"\3\'\3\'\3(\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\6.\u0120\n.\r.\16."+
		"\u0121\3.\3.\3/\3/\3/\2\2\60\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61"+
		"\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\2\3\2"+
		"\n\3\2\61\61\3\2\f\f\3\2\62;\3\2$$\7\2C\\aac|\u00d3\u00d3\u00f3\u00f3"+
		"\7\2\62;C\\c|\u00d3\u00d3\u00f3\u00f3\5\2\13\f\17\17\"\"\t\2\"#%%--/\60"+
		"<<BB]_\2\u012f\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3"+
		"\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2"+
		"\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3"+
		"\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2"+
		"\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\2"+
		"9\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3"+
		"\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2"+
		"\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\3_\3\2\2\2\5"+
		"g\3\2\2\2\7m\3\2\2\2\tq\3\2\2\2\13x\3\2\2\2\r\u0080\3\2\2\2\17\u0085\3"+
		"\2\2\2\21\u008d\3\2\2\2\23\u0091\3\2\2\2\25\u0094\3\2\2\2\27\u009d\3\2"+
		"\2\2\31\u00a1\3\2\2\2\33\u00a4\3\2\2\2\35\u00b1\3\2\2\2\37\u00bc\3\2\2"+
		"\2!\u00c1\3\2\2\2#\u00cb\3\2\2\2%\u00cf\3\2\2\2\'\u00e1\3\2\2\2)\u00e3"+
		"\3\2\2\2+\u00ea\3\2\2\2-\u00ec\3\2\2\2/\u00ee\3\2\2\2\61\u00f0\3\2\2\2"+
		"\63\u00f2\3\2\2\2\65\u00f4\3\2\2\2\67\u00f6\3\2\2\29\u00f8\3\2\2\2;\u00fa"+
		"\3\2\2\2=\u00fc\3\2\2\2?\u00fe\3\2\2\2A\u0100\3\2\2\2C\u0102\3\2\2\2E"+
		"\u0104\3\2\2\2G\u0107\3\2\2\2I\u010a\3\2\2\2K\u010c\3\2\2\2M\u010e\3\2"+
		"\2\2O\u0111\3\2\2\2Q\u0114\3\2\2\2S\u0116\3\2\2\2U\u0118\3\2\2\2W\u011a"+
		"\3\2\2\2Y\u011c\3\2\2\2[\u011f\3\2\2\2]\u0125\3\2\2\2_`\7r\2\2`a\7t\2"+
		"\2ab\7k\2\2bc\7p\2\2cd\7v\2\2de\7n\2\2ef\7p\2\2f\4\3\2\2\2gh\7r\2\2hi"+
		"\7t\2\2ij\7k\2\2jk\7p\2\2kl\7v\2\2l\6\3\2\2\2mn\7n\2\2no\7g\2\2op\7v\2"+
		"\2p\b\3\2\2\2qr\7U\2\2rs\7v\2\2st\7t\2\2tu\7k\2\2uv\7p\2\2vw\7i\2\2w\n"+
		"\3\2\2\2xy\7K\2\2yz\7p\2\2z{\7v\2\2{|\7g\2\2|}\7i\2\2}~\7g\2\2~\177\7"+
		"t\2\2\177\f\3\2\2\2\u0080\u0081\7T\2\2\u0081\u0082\7g\2\2\u0082\u0083"+
		"\7c\2\2\u0083\u0084\7n\2\2\u0084\16\3\2\2\2\u0085\u0086\7D\2\2\u0086\u0087"+
		"\7q\2\2\u0087\u0088\7q\2\2\u0088\u0089\7n\2\2\u0089\u008a\7g\2\2\u008a"+
		"\u008b\7c\2\2\u008b\u008c\7p\2\2\u008c\20\3\2\2\2\u008d\u008e\7u\2\2\u008e"+
		"\u008f\7v\2\2\u008f\u0090\7t\2\2\u0090\22\3\2\2\2\u0091\u0092\7K\2\2\u0092"+
		"\u0093\7h\2\2\u0093\24\3\2\2\2\u0094\u0095\7g\2\2\u0095\u0096\7p\2\2\u0096"+
		"\u0097\7v\2\2\u0097\u0098\7q\2\2\u0098\u0099\7p\2\2\u0099\u009a\7e\2\2"+
		"\u009a\u009b\7g\2\2\u009b\u009c\7u\2\2\u009c\26\3\2\2\2\u009d\u009e\7"+
		"C\2\2\u009e\u009f\7P\2\2\u009f\u00a0\7F\2\2\u00a0\30\3\2\2\2\u00a1\u00a2"+
		"\7Q\2\2\u00a2\u00a3\7T\2\2\u00a3\32\3\2\2\2\u00a4\u00a5\7\61\2\2\u00a5"+
		"\u00a6\7,\2\2\u00a6\u00a8\3\2\2\2\u00a7\u00a9\n\2\2\2\u00a8\u00a7\3\2"+
		"\2\2\u00a9\u00aa\3\2\2\2\u00aa\u00a8\3\2\2\2\u00aa\u00ab\3\2\2\2\u00ab"+
		"\u00ac\3\2\2\2\u00ac\u00ad\7,\2\2\u00ad\u00ae\7\61\2\2\u00ae\u00af\3\2"+
		"\2\2\u00af\u00b0\b\16\2\2\u00b0\34\3\2\2\2\u00b1\u00b2\7\61\2\2\u00b2"+
		"\u00b3\7\61\2\2\u00b3\u00b5\3\2\2\2\u00b4\u00b6\n\3\2\2\u00b5\u00b4\3"+
		"\2\2\2\u00b6\u00b7\3\2\2\2\u00b7\u00b5\3\2\2\2\u00b7\u00b8\3\2\2\2\u00b8"+
		"\u00b9\3\2\2\2\u00b9\u00ba\b\17\2\2\u00ba\36\3\2\2\2\u00bb\u00bd\t\4\2"+
		"\2\u00bc\u00bb\3\2\2\2\u00bd\u00be\3\2\2\2\u00be\u00bc\3\2\2\2\u00be\u00bf"+
		"\3\2\2\2\u00bf \3\2\2\2\u00c0\u00c2\t\4\2\2\u00c1\u00c0\3\2\2\2\u00c2"+
		"\u00c3\3\2\2\2\u00c3\u00c1\3\2\2\2\u00c3\u00c4\3\2\2\2\u00c4\u00c5\3\2"+
		"\2\2\u00c5\u00c7\7\60\2\2\u00c6\u00c8\t\4\2\2\u00c7\u00c6\3\2\2\2\u00c8"+
		"\u00c9\3\2\2\2\u00c9\u00c7\3\2\2\2\u00c9\u00ca\3\2\2\2\u00ca\"\3\2\2\2"+
		"\u00cb\u00cc\7)\2\2\u00cc\u00cd\n\5\2\2\u00cd\u00ce\7)\2\2\u00ce$\3\2"+
		"\2\2\u00cf\u00d3\7$\2\2\u00d0\u00d2\n\5\2\2\u00d1\u00d0\3\2\2\2\u00d2"+
		"\u00d5\3\2\2\2\u00d3\u00d1\3\2\2\2\u00d3\u00d4\3\2\2\2\u00d4\u00d6\3\2"+
		"\2\2\u00d5\u00d3\3\2\2\2\u00d6\u00d7\7$\2\2\u00d7&\3\2\2\2\u00d8\u00d9"+
		"\7v\2\2\u00d9\u00da\7t\2\2\u00da\u00db\7w\2\2\u00db\u00e2\7g\2\2\u00dc"+
		"\u00dd\7h\2\2\u00dd\u00de\7c\2\2\u00de\u00df\7n\2\2\u00df\u00e0\7u\2\2"+
		"\u00e0\u00e2\7g\2\2\u00e1\u00d8\3\2\2\2\u00e1\u00dc\3\2\2\2\u00e2(\3\2"+
		"\2\2\u00e3\u00e7\t\6\2\2\u00e4\u00e6\t\7\2\2\u00e5\u00e4\3\2\2\2\u00e6"+
		"\u00e9\3\2\2\2\u00e7\u00e5\3\2\2\2\u00e7\u00e8\3\2\2\2\u00e8*\3\2\2\2"+
		"\u00e9\u00e7\3\2\2\2\u00ea\u00eb\7?\2\2\u00eb,\3\2\2\2\u00ec\u00ed\7\60"+
		"\2\2\u00ed.\3\2\2\2\u00ee\u00ef\7=\2\2\u00ef\60\3\2\2\2\u00f0\u00f1\7"+
		".\2\2\u00f1\62\3\2\2\2\u00f2\u00f3\7<\2\2\u00f3\64\3\2\2\2\u00f4\u00f5"+
		"\7#\2\2\u00f5\66\3\2\2\2\u00f6\u00f7\7(\2\2\u00f78\3\2\2\2\u00f8\u00f9"+
		"\7*\2\2\u00f9:\3\2\2\2\u00fa\u00fb\7+\2\2\u00fb<\3\2\2\2\u00fc\u00fd\7"+
		"}\2\2\u00fd>\3\2\2\2\u00fe\u00ff\7\177\2\2\u00ff@\3\2\2\2\u0100\u0101"+
		"\7]\2\2\u0101B\3\2\2\2\u0102\u0103\7_\2\2\u0103D\3\2\2\2\u0104\u0105\7"+
		"?\2\2\u0105\u0106\7?\2\2\u0106F\3\2\2\2\u0107\u0108\7#\2\2\u0108\u0109"+
		"\7?\2\2\u0109H\3\2\2\2\u010a\u010b\7@\2\2\u010bJ\3\2\2\2\u010c\u010d\7"+
		">\2\2\u010dL\3\2\2\2\u010e\u010f\7@\2\2\u010f\u0110\7?\2\2\u0110N\3\2"+
		"\2\2\u0111\u0112\7>\2\2\u0112\u0113\7?\2\2\u0113P\3\2\2\2\u0114\u0115"+
		"\7,\2\2\u0115R\3\2\2\2\u0116\u0117\7\61\2\2\u0117T\3\2\2\2\u0118\u0119"+
		"\7-\2\2\u0119V\3\2\2\2\u011a\u011b\7/\2\2\u011bX\3\2\2\2\u011c\u011d\7"+
		"\'\2\2\u011dZ\3\2\2\2\u011e\u0120\t\b\2\2\u011f\u011e\3\2\2\2\u0120\u0121"+
		"\3\2\2\2\u0121\u011f\3\2\2\2\u0121\u0122\3\2\2\2\u0122\u0123\3\2\2\2\u0123"+
		"\u0124\b.\2\2\u0124\\\3\2\2\2\u0125\u0126\7^\2\2\u0126\u0127\t\t\2\2\u0127"+
		"^\3\2\2\2\f\2\u00aa\u00b7\u00be\u00c3\u00c9\u00d3\u00e1\u00e7\u0121\3"+
		"\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}